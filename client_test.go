package peloton

import (
	"strings"
	"testing"
)

func TestNewPelotonClientURL(t *testing.T) {
	result := NewPelotonClient("", "")
	if result.client.HostURL != "https://api.onepeloton.com" {
		t.Errorf("API URL should be set, got %v", result.client.HostURL)
	}
}

func TestNewPelotonClientResponses(t *testing.T) {
	result := NewPelotonClient("", "")
	resp, _ := result.client.R().
		SetHeader("Accept", "application/json").
		Get("api/instructor")
	if resp.StatusCode() != 200 {
		t.Errorf("Client should receive http 200 code from unauth endpoint, got %v", resp.StatusCode())
	}
	if strings.Join(resp.Header()["Content-Type"], "") != "application/json" {
		t.Errorf("Expected json content type got %v", resp.Header()["Content-Type"])
	}
}

func TestPelotonClient_Instructors(t *testing.T) {
	client := NewPelotonClient("", "")
	m, _ := client.Instructors()
	if m.Count <= 0 {
		t.Errorf("Should have more than 0 instructors, got %v", m.Count)
	}
	if m.Total <= 0 {
		t.Errorf("Should have more than 0 instructors, got %v", m.Total)
	}
}
