// Copyright (c) 2020 Chris McFee (cmcfee@kent.edu), All rights reserved.
// Source code and usage is governed by an Apache style license
// that can be found in the LICENSE file.

// Package peloton provides a simple http rest client for the Peloton APIs
package peloton

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/mickfeech/peloton/models"
	"os"
)

// Default URL for all API requests
const defaultBaseURL = "https://api.onepeloton.com"

// PelotonClient struct is used to create a resty client client
type PelotonClient struct {
	client *resty.Client
}

// Error struct to return messages from resty client
type Error struct {
	Code    string `json:"error_code,omitempty"`
	Message string `json:"error_message,omitempty"`
}

// NewPelotonClient method creates a new client using the resty package
func NewPelotonClient(username string, password string) *PelotonClient {
	client := resty.New()
	client.SetDebug(false)
	apiURL := os.Getenv("API_ADDR")
	if apiURL == "" {
		apiURL = defaultBaseURL
	}
	client.SetHostURL(apiURL)
	client.SetError(&Error{})
	auth := fmt.Sprintf("{\"username_or_email\": \"%v\", \"password\": \"%v\"}", username, password)
	authData := []byte(auth)
	client.R().
		SetBody(authData).
		Post("/auth/login")
	return &PelotonClient{client: client}
}

// Me method creates a request to retrieve data about the current user
func (c *PelotonClient) Me() (*models.User, error) {
	resp, err := c.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&models.User{}).
		Get("/api/me")
	if err != nil {
		return nil, err
	}
	//return resp.String(), err
	return resp.Result().(*models.User), nil
}

// Instructors method creates a request to retrieve data about the instructors
func (c *PelotonClient) Instructors() (*models.Instructors, error) {
	resp, err := c.client.R().
		SetHeader("Accept", "application/json").
		SetResult(&models.Instructors{}).
		Get("/api/instructor")
	if err != nil {
		return nil, err
	}
	return resp.Result().(*models.Instructors), err
}
