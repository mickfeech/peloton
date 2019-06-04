package peloton

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
)

const BaseAddress = "https://api.onepeloton.com/api"

type Client struct {
	Client  *http.Client
	baseURL string
}

func NewClient() *Client {
	return &Client{
		Client:  http.DefaultClient,
		baseURL: BaseAddress,
	}
}

func (c *Client) get(url string, result interface{}) error {
	for {
		transCfg := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		c.Client = &http.Client{Transport: transCfg}
		resp, err := c.Client.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(result)
		if err != nil {
			return err
		}
		break
	}
	return nil
}
