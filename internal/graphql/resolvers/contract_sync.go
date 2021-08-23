// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"bytes"
	"context"
	"encoding/json"
	"fantom-api-graphql/internal/types"
	"net/http"
	"sync"
	"time"
)

const (
	// contractSyncMutationQuery represents the mutation GraphQL query used
	// to synchronize contract validation with API peers.
	contractSyncMutationQuery = "mutation($sc:ContractValidationInput!) { validateContract(contract: $sc) { validated } }"

	// contractSyncCallTimeout represents a time out value used for contract
	// syncing GraphQL calls.
	contractSyncCallTimeout = 60 * time.Second
)

// getContractSyncInput prepares input structure used for contract syncing
// across peer API points.
func contractSyncInput(con *types.Contract) ContractValidationInput {
	// prep the validation input to be synced
	var cInput = ContractValidationInput{
		Address:      con.Address,
		Name:         &con.Name,
		SourceCode:   con.SourceCode,
		OptimizeRuns: con.OptimizeRuns,
		Optimized:    con.IsOptimized,
	}

	// transfer compiler version info, if any
	if 0 < len(con.Version) {
		cInput.Version = &con.Version
	}

	// transfer support contact info, if any
	if 0 < len(con.SupportContact) {
		cInput.SupportContact = &con.SupportContact
	}

	// transfer contact open source license, if any
	if 0 < len(con.License) {
		cInput.License = &con.License
	}

	return cInput
}

// constructMutation creates the GraphQL mutation query string
// for the contract provided.
func constructMutationPayload(con *types.Contract) (bytes.Buffer, error) {
	// prepare the payload
	payload := struct {
		Query     string                 `json:"query"`
		Variables map[string]interface{} `json:"variables,omitempty"`
	}{
		Query: contractSyncMutationQuery,
		Variables: map[string]interface{}{
			"sc": contractSyncInput(con),
		},
	}

	// encode the mutation into the output buffer
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(payload)
	if err != nil {
		return buf, err
	}

	return buf, nil
}

// SyncContract synchronizes contract across all the peers in the API network.
func (rs *rootResolver) syncContract(con types.Contract) {
	// no peers to sync against
	if len(cfg.Server.Peers) <= 0 {
		log.Debugf("no peers for contract validation syncing")
		return
	}

	// construct the payload
	payload, err := constructMutationPayload(&con)
	if err != nil {
		log.Errorf("can not construct the sync payload; %s", err.Error())
		return
	}

	// prep wait group to sync all routines
	var wg sync.WaitGroup

	// loop over the peers and sync each of them
	for _, peer := range cfg.Server.Peers {
		// add this sync to the wait group
		wg.Add(1)

		// run the sync
		go syncContractToPeer(&payload, peer, cfg.Server.DomainAddress, &wg)
	}

	// wait for all the sync to finish
	log.Debugf("waiting for validation syncing to finish")
	wg.Wait()

	// inform we are done here
	log.Debugf("validation syncing finished")
}

// syncContractToPeer performs the syncing call for the contract validation.
func syncContractToPeer(payload *bytes.Buffer, peer string, origin string, wg *sync.WaitGroup) {
	// log action
	log.Debugf("syncing contract validation to %s from %s", peer, origin)

	// make a context with predefined timeout, we don't use the cancel func callback
	ctx, cancel := context.WithTimeout(context.Background(), contractSyncCallTimeout)

	// don't forget to sign off after we are done
	defer func() {
		// log finish
		cancel()
		log.Noticef("syncing %s finished", peer)

		// signal to wait group we are done
		wg.Done()
	}()

	// create the request
	req, err := http.NewRequestWithContext(ctx, "POST", peer, payload)
	if err != nil {
		log.Errorf("can not create new POST request for %s peer", peer)
		return
	}

	// set headers so we can pass the payload correctly
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", origin)

	// make the client and send the request
	client := &http.Client{}

	// fire the request
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("can not finalize syncing request for %s peer; %s", peer, err.Error())
		return
	}

	// log error code response
	if 200 != resp.StatusCode {
		log.Errorf("syncing request to %s has been rejected with code %d", peer, resp.StatusCode)
		return
	}

	// success
	log.Debugf("syncing request to %s finished with success", peer)
}
