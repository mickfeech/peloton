package peloton

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
)

// BaseAddress is the api endpoint for peloton
const BaseAddress = "https://api.onepeloton.com/api"

// Client is a client for working with the peloton web api
type Client struct {
	Client  *http.Client
	baseURL string
}

// NewClient initializes a new HTTP client to interact with the web api
func NewClient() *Client {
	return &Client{
		Client:  http.DefaultClient,
		baseURL: BaseAddress,
	}
}

// Get executes a HTTP request to the web api
func (c *Client) Get(url string, result interface{}) error {
	for {
		transCfg := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		c.Client = &http.Client{Transport: transCfg}
		resp, err := c.Client.Get(url)
		defer resp.Body.Close()
		if err != nil {
			return err
		}
		err = json.NewDecoder(resp.Body).Decode(result)
		if err != nil {
			return err
		}
		break
	}
	return nil
}
