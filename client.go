package peloton

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"time"
)

// BaseAddress is the api endpoint for peloton
const BaseAddress = "https://api.onepeloton.com/api"

// Client is a client for working with the peloton web api
type Client struct {
	Client  *http.Client
	baseURL string
	jar     http.CookieJar
}

// NewClient initializes a new HTTP client to interact with the web api
func NewClient() *Client {
	j, _ := cookiejar.New(nil)
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	jClient := &http.Client{Jar: j, Transport: transCfg, Timeout: time.Second * 10}

	return &Client{
		Client:  jClient,
		baseURL: BaseAddress,
	}
}

// Get executes a HTTP request to the web api
func (c *Client) Get(url string, result interface{}) error {
	resp, err := c.Client.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

// GetCookie executes a request to set the header to maintain session with the api
func (c *Client) GetCookie(apiUrl string, data []byte) []*http.Cookie {
	req, _ := http.NewRequest("POST", apiUrl, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := c.Client.Do(req)
	defer resp.Body.Close()
	return resp.Cookies()
}
