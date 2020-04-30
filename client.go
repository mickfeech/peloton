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
	Client   *resty.Client
	Username string
	Password string
}

func NewApiClient(args ...interface{}) (*ApiClient, error) {
	client := resty.New()
	client.SetDebug(false)
	apiURL := os.Getenv("API_ADDR")
	if apiURL == "" {
		apiURL = defaultBaseURL
	}
	var username string
	var password string

	if len(args) == 2 {
		for i, arg := range args {
			switch i {
			case 0: //username
				username, _ = arg.(string)
			case 1: //password
				password, _ = arg.(string)
			}
		}
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
	return &ApiClient{Client: client, Username: username, Password: password}, err
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
func (c *ApiClient) GetInstructors() (*models.Instructors, error) {
	resp, _ := c.Client.R().
		SetHeader("Accept", "application/json").
		SetResult(&models.Instructors{}).
		Get("/api/instructor")
	return resp.Result().(*models.Instructors), nil
}

// Workouts method creates a request to retrieve workout data by userid
func (c *ApiClient) GetUserWorkouts(userid string) (*models.Workouts, error) {
	userWorkoutUrl := fmt.Sprintf("/api/user/%s/workouts", userid)
	resp, _ := c.Client.R().
		SetHeader("Accept", "application/json").
		SetResult(&models.Workouts{}).
		Get(userWorkoutUrl)
	return resp.Result().(*models.Workouts), nil
}

// Get specific workout details
func (c *ApiClient) GetWorkoutDetails(workoutId string, interval string) (*models.WorkOutDetails, error) {
	workoutDetailUrl := fmt.Sprintf("/api/workout/%s/performance_graph?every_n=%s", workoutId, interval)
	resp, _ := c.Client.R().
		SetHeader("Accept", "application/json").
		SetResult(&models.WorkOutDetails{}).
		Get(workoutDetailUrl)
	return resp.Result().(*models.WorkOutDetails), nil
}
