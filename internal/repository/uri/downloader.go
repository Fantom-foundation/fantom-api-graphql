// Package uri provides tools for downloading data from URI - especially
// NFT tokens JSON Metadata and related images using IPFS or HTTP.
package uri

import (
	"encoding/base64"
	"errors"
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/logger"
	"fmt"
	ipfsapi "github.com/ipfs/go-ipfs-api"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ipfsRequestTimeout represents the timeout applied to IPFS requests.
const ipfsRequestTimeout = 30 * time.Second

type Downloader struct {
	ipfsShell        *ipfsapi.Shell
	skipHttpGateways bool
	gateway          string
	gatewayBearer    string

	// log represents the logger to be used by the repository.
	log logger.Logger
}

// New provides new Downloader instance.
func New(cfg *config.Config, l logger.Logger) *Downloader {
	d := &Downloader{
		ipfsShell:        ipfsapi.NewShell(cfg.Ipfs.Url),
		skipHttpGateways: cfg.Ipfs.SkipHttpGateways,
		gateway:          cfg.Ipfs.Gateway,
		gatewayBearer:    cfg.Ipfs.GatewayBearer,
		log:              l,
	}
	d.ipfsShell.SetTimeout(ipfsRequestTimeout)
	return d
}

// unescapeIPFSUri decodes URL escaped address.
func unescapeIPFSUri(uri string) string {
	iUri, err := url.QueryUnescape(uri)
	if err != nil {
		return "/ipfs/" + uri
	}
	return "/ipfs/" + iUri
}

// GetIpfsUri try to obtain IPFS URI from the URI - returns empty string for non-ipfs uri
// This function is responsible for IPFS URI detection, unification and for conversion
// of known IPFS HTTP gateways URI into IPFS URI.
func (d *Downloader) GetIpfsUri(uri string) string {
	if strings.HasPrefix(uri, "/ipfs/") {
		return uri
	}
	if strings.HasPrefix(uri, "ipfs://") {
		return "/ipfs/" + uri[7:]
	}

	if d.skipHttpGateways {
		if strings.HasPrefix(uri, "https://gateway.pinata.cloud/ipfs/") {
			return unescapeIPFSUri(uri[34:])
		}
		if strings.HasPrefix(uri, "https://ipfs.io/ipfs/") {
			return unescapeIPFSUri(uri[21:])
		}
		if idx := strings.Index(uri, ".mypinata.cloud/ipfs/"); idx > 8 && idx < 30 {
			return unescapeIPFSUri(uri[idx+21:])
		}
	}
	return ""
}

// GetFromIpfs downloads the file from IPFS (URI is expected in "/ipfs/{CID}" form).
func (d *Downloader) GetFromIpfs(uri string) (data []byte, mimetype string, err error) {
	if d.gateway != "" {
		return d.GetFromIpfsGateway(uri)
	}

	reader, err := d.ipfsShell.Cat(uri)
	if err != nil {
		return nil, "", err
	}

	defer func(r io.ReadCloser) {
		if e := r.Close(); e != nil {
			d.log.Errorf("could not close IPFS shell reader; %s", e.Error())
		}
	}(reader)

	out, err := io.ReadAll(reader)
	if err != nil {
		return nil, "", err
	}
	return out, "", nil
}

// GetFromIpfsGateway downloads the file from IPFS HTTP gateway.
func (d *Downloader) GetFromIpfsGateway(uri string) (data []byte, mimetype string, err error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", d.gateway+uri, nil)
	if err != nil {
		return nil, "", err
	}

	// add access authorization header with pre-shared bearer token
	if d.gatewayBearer != "" {
		req.Header.Set("Authorization", "Bearer "+d.gatewayBearer)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}

	defer func(r io.ReadCloser) {
		if e := r.Close(); e != nil {
			d.log.Errorf("could not close IPFS gateway response reader; %s", e.Error())
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		return nil, "", fmt.Errorf("HTTP gateway returned %s", resp.Status)
	}

	out, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	mimetype = resp.Header.Get("Content-Type")
	return out, mimetype, nil
}

// GetFromHttp downloads the file from HTTP.
func (d *Downloader) GetFromHttp(uri string) (data []byte, mimetype string, err error) {
	client := http.Client{
		Timeout: 1 * time.Minute,
	}

	// call for image
	resp, err := client.Get(uri)
	if err != nil {
		return nil, "", fmt.Errorf("request failed; %s", err.Error())
	}

	// make sure to always close the body reader
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			d.log.Errorf("error closing HTTP body for %s; %s", uri, err.Error())
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		return nil, "", fmt.Errorf("HTTP server returned %s", resp.Status)
	}

	out, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("can not read response; %s", err.Error())
	}

	mimetype = resp.Header.Get("Content-Type")
	return out, mimetype, nil
}

// GetFromDataUri obtains the file encoded in "data:" URI.
func (d *Downloader) GetFromDataUri(uri string) (data []byte, mimetype string, err error) {
	parts := strings.Split(uri, ",")
	if len(parts) < 2 {
		return nil, "", errors.New("Invalid data uri - no comma: " + uri)
	}
	mimetype = strings.Split(parts[0][5:], ";")[0]

	out, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, "", err
	}
	return out, mimetype, nil
}
