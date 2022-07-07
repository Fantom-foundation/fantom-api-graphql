package repository

import (
	"fantom-api-graphql/internal/types"
	"fmt"
	neturl "net/url"
	"strings"
)

// GetTokenJsonMetadata provides decoded JSON metadata for the given token metadata URI.
func (p *proxy) GetTokenJsonMetadata(uri string) (*types.NftMetadata, error) {
	data, err := p.GetRawTokenJsonMetadata(uri)
	if err != nil {
		return nil, err
	}
	jsonMeta, err := types.DecodeNftMetadata(data)
	if err != nil {
		return nil, fmt.Errorf("unable to decode json; %s", err)
	}
	return jsonMeta, err
}

// GetRawTokenJsonMetadata provides decoded JSON metadata for the given token metadata URI.
func (p *proxy) GetRawTokenJsonMetadata(uri string) ([]byte, error) {
	// make sure to do this only once, if parallel requests were fired
	var key strings.Builder
	key.WriteString("GetRawTokenJsonMetadata")
	key.WriteString(uri)

	data, err, _ := p.apiRequestGroup.Do(key.String(), func() (interface{}, error) {
		data, _, err := p.getFileFromUri(uri)
		if err != nil {
			return nil, fmt.Errorf("unable to download json; %s", err)
		}
		return data, err
	})
	if err != nil {
		return nil, err
	}
	return data.([]byte), err
}

// getCidFromIpfsUri extracts IPFS CID from the given URI.
func getCidFromIpfsUri(uri string) string {
	cid := uri[6:] // skip /ipfs/
	slashIdx := strings.Index(cid, "/")
	if slashIdx != -1 {
		cid = cid[0:slashIdx]
	}
	return cid
}

// getFileFromUri resolves the URI and download file from the URI using appropriate protocol
func (p *proxy) getFileFromUri(uri string) (data []byte, mimetype string, err error) {
	// the URI contains the data directly as BASE64 encoded data stream
	if strings.HasPrefix(uri, "data:") {
		return p.uri.GetFromDataUri(uri)
	}

	// serve IPFS (or IPFS gateway) URIs
	if ipfsUri := p.uri.GetIpfsUri(uri); ipfsUri != "" {
		// try the local cache first
		cachedContent := p.filecache.PullIpfsFile(getCidFromIpfsUri(ipfsUri))
		if cachedContent != nil {
			return cachedContent, "", nil
		}

		// extract data from IPFS, if possible
		data, mimetype, err = p.uri.GetFromIpfs(ipfsUri)
		if err == nil || strings.HasPrefix(uri, "ipfs:") || strings.HasPrefix(uri, "/ipfs/") {
			return data, mimetype, err
		} else {
			// if err, but it is http ipfs proxy uri, only log and continue to process as a http uri
			log.Warningf("failed to download file %s from IPFS, will try http; %s", ipfsUri, err)
		}
	}

	// try to decode the URI to validate the format
	url, err := neturl.Parse(uri)
	if err != nil {
		return nil, "", fmt.Errorf("invalid URI format at %s; %s", uri, err.Error())
	}

	// check the schema for known access methods
	if url.Scheme == "http" || url.Scheme == "https" {
		return p.uri.GetFromHttp(uri)
	}

	return nil, "", fmt.Errorf("invalid URI scheme; file not available at %s", uri)
}
