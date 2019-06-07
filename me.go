package peloton

// Instructor contains basic information about instructors
type User struct {
	ID             string     `json:"id"`
	FirstName      string     `json:"first_name"`
	LastName       string     `json:"last_name"`
	Username       string     `json:"username"`
	Ftp            int        `json:"cycling_workout_ftp"`
	CustomMaxHr    int        `json:"customized_max_heart_rate"`
	TotalWorkOuts  int        `json:"total_workouts"`
	DefaultMaxHr   int        `json:"default_max_heart_rate"`
	EstimatedFtp   int        `json:"estimated_cycling_ftp"`
	DefaultHrZones []float64  `json:"default_heart_rate_zones"`
	WorkOutCounts  []WorkOuts `json:"workout_counts"`
}

// WorkOuts
type WorkOuts struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
	Slug  string `json:"slug"`
}

// GetMyInfo gets current user info
func (c *Client) GetMyInfo() (*User, error) {
	pelotonURL := c.baseURL + "/me"
	var t User
	err := c.Get(pelotonURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
