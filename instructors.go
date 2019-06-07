package peloton

// InstructorsResponse contains data element from initial call to the api
type InstructorsResponse struct {
	Data []Instructor `json:"data"`
}

// Instructor contains basic information about instructors
type Instructor struct {
	ID              string   `json:"id"`
	FirstName       string   `json:"first_name"`
	LastName        string   `json:"last_name"`
	Disciplines     []string `json:"fitness_disciplines"`
	Username        string   `json:"username"`
	SpotifyPlaylist string   `json:"spotify_playlist_uri"`
}

// GetInstructors gets instructor information
func (c *Client) GetInstructors() (*InstructorsResponse, error) {
	pelotonURL := c.baseURL + "/instructor"
	var t InstructorsResponse
	err := c.Get(pelotonURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
