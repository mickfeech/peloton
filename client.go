package peloton

import (
	"fmt"
	"github.com/go-resty/resty"
	"os"
)

const defaultBaseURL = "https://api.onepeloton.com"

type PelotonClient struct {
	client *resty.Client
}

type Error struct {
	Code    string `json:"error_code,omitempty"`
	Message string `json:"error_message,omitempty"`
}

type User struct {
	ID             string    `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Username       string    `json:"username"`
	Ftp            int       `json:"cycling_workout_ftp"`
	CustomMaxHr    int       `json:"customized_max_heart_rate"`
	TotalWorkOuts  int       `json:"total_workouts"`
	DefaultMaxHr   int       `json:"default_max_heart_rate"`
	EstimatedFtp   int       `json:"estimated_cycling_ftp"`
	DefaultHrZones []float64 `json:"default_heart_rate_zones"`
	//WorkOutCounts  []WorkOuts `json:"workout_counts"`
}

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

func (c *PelotonClient) Me() (string, error) {
	resp, err := c.client.R().
		SetHeader("Accept", "application/json").
		Get("/api/me")
	if err != nil {
		return "", err
	}
	return resp.String(), err
}

func (c *PelotonClient) Instructors() (string, error) {
	resp, err := c.client.R().
		SetHeader("Accept", "application/json").
		Get("/api/instructor")
	if err != nil {
		return "", err
	}
	return resp.String(), err
}
