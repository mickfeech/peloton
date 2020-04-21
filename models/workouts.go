package models

type Workouts struct {
	Count int `json:"count"`
	Data  []WorkoutData
}

type WorkoutData struct {
	FitnessDiscipline string `json:"fitness_discipline"`
	Created           int    `json:"created_at"`
	EndTime           int    `json:"end_time"`
	ID                string `json:"id"`
	StartTime         int    `json:"start_time"`
	UserID            string `json:"user_id"`
	Status            string `json:"status"`
	PersonalRecord    bool   `json:"is_total_work_personal_record"`
}
