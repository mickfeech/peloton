package peloton

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/mickfeech/peloton/models"
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

//func (c *PelotonClient) Me() (string, error) {
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
