package models

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
	Birthday       int       `json:"birthday"`
	Location       string    `json:"location"`
	Height         float64   `json:"height"`
	Weight         float64   `json:"weight"`
	Gender         string    `json:"gender"`
	//WorkOutCounts  []WorkOuts `json:"workout_counts"`
}
