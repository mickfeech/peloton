package models

type Instructors struct {
	Total       int          `json:"total"`
	Count       int          `json:"count"`
	Instructors []Instructor `json:"data"`
}

type Instructor struct {
	ID          string   `json:"id"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	Disciplines []string `json:"fitness_disciplines"`
	Username    string   `json:"username"`
}
