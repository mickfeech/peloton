// Copyright (c) 2020 Chris McFee (cmcfee@kent.edu), All rights reserved.
// Source code and usage is governed by an Apache style license
// that can be found in the LICENSE file.

// Package peloton provides a simple http rest client for the Peloton APIs
package peloton

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
	"os"
)

// Default URL for all API requests
const defaultBaseURL = "https://api.onepeloton.com"

// APIClient to allow updates to the client instance
type ApiClient struct {
	Client   *resty.Client
	Username string
	Password string
	Cookie   *http.Cookie
}

// instantiate a new instance of the API client
func NewApiClient(args ...interface{}) (*ApiClient, error) {
	client := resty.New()
	apiURL := os.Getenv("API_ADDR")
	if apiURL == "" {
		apiURL = defaultBaseURL
	}
	var username string
	var password string
	if len(args) == 2 {
		for i, arg := range args {
			switch i {
			case 0:
				username = arg.(string)
			case 1:
				password = arg.(string)
			}
		}
	}
	client.SetHostURL(apiURL)
	cookie, _ := GetInitialAuthCookie(client, username, password)
	return &ApiClient{Client: client, Username: username, Password: password, Cookie: cookie}, nil
}

// GetAuthCookie method to set the peloton_session_id cookie
func GetInitialAuthCookie(c *resty.Client, username string, password string) (*http.Cookie, error) {
	authData := []byte(fmt.Sprintf("{\"username_or_email\": \"%v\", \"password\": \"%v\"}", username, password))
	resp, err := c.R().
		SetBody(authData).
		Post("/auth/login")
	if err != nil || resp.IsError() {
		if err != nil {
			log.Fatal(err)
		}
	}
	var authCookie *http.Cookie
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "peloton_session_id" {
			authCookie = cookie
		}
	}
	return authCookie, nil
}

// Update AuthCookie from client
func (c *ApiClient) UpdateAuthCookie() (*http.Cookie, error) {
	cookie, _ := GetInitialAuthCookie(c.Client, c.Username, c.Password)
	c.Cookie = cookie
	return cookie, nil
}
