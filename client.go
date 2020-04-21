// Copyright (c) 2020 Chris McFee (cmcfee@kent.edu), All rights reserved.
// Source code and usage is governed by an Apache style license
// that can be found in the LICENSE file.

// Package peloton provides a simple http rest client for the Peloton APIs
package peloton

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/mickfeech/peloton/models"
	"log"
	"os"
)

// Default URL for all API requests
const defaultBaseURL = "https://api.onepeloton.com"

type ApiClient struct {
	Client *resty.Client
}

func NewApiClient(username string, password string) (*ApiClient, error) {
	client := resty.New()
	client.SetDebug(false)
	apiURL := os.Getenv("API_ADDR")
	if apiURL == "" {
		apiURL = defaultBaseURL
	}
	client.SetHostURL(apiURL)
	auth := fmt.Sprintf("{\"username_or_email\": \"%v\", \"password\": \"%v\"}", username, password)
	authData := []byte(auth)
	resp, err := client.R().
		SetBody(authData).
		Post("/auth/login")
	if err != nil || resp.IsError() {
		if err != nil {
			log.Fatal(err)
		}
	}
	return &ApiClient{Client: client}, err
}

// Me method creates a request to retrieve data about the current user
func (c *ApiClient) Me() (*models.User, error) {
	resp, _ := c.Client.R().
		SetHeader("Accept", "application/json").
		SetResult(&models.User{}).
		Get("/api/me")
	return resp.Result().(*models.User), nil
}

// Instructors method creates a request to retrieve data about the instructors
func (c *ApiClient) Instructors() (*models.Instructors, error) {
	resp, _ := c.Client.R().
		SetHeader("Accept", "application/json").
		SetResult(&models.Instructors{}).
		Get("/api/instructor")
	return resp.Result().(*models.Instructors), nil
}

// Workouts method creates a request to retrieve workout data by userid
func (c *ApiClient) Workouts(userid string) (*models.Workouts, error) {
	workoutUrl := fmt.Sprintf("/api/user/%s/workouts", userid)
	resp, _ := c.Client.R().
		SetHeader("Accept", "application/json").
		SetResult(&models.Workouts{}).
		Get(workoutUrl)
	return resp.Result().(*models.Workouts), nil
}
