package peloton

type InstructorsResponse struct {
	Data []Instructor `json:"data"`
}

type Instructor struct {
	ID              string   `json:"id"`
	FirstName       string   `json:"first_name"`
	LastName        string   `json:"last_name"`
	Disciplines     []string `json:"fitness_disciplines"`
	Username        string   `json:"username"`
	SpotifyPlaylist string   `json:"spotify_playlist_uri"`
}

func (c *Client) GetInstructors() (*InstructorsResponse, error) {
	pelotonURL := c.baseURL + "/instructor"
	var t InstructorsResponse
	err := c.get(pelotonURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
