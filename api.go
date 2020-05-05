package peloton

import (
	"fmt"
	"github.com/mickfeech/peloton/models"
)

// Instructors method creates a request to retrieve data about the instructors
func (c *ApiClient) GetInstructors() (*models.Instructors, error) {
	resp, _ := c.Client.R().
		SetHeader("Accept", "application/json").
		SetResult(&models.Instructors{}).
		Get("/api/instructor")
	return resp.Result().(*models.Instructors), nil
}

// GetSchedule
func (c *ApiClient) GetSchedule() (*models.Schedule, error) {
	resp, _ := c.Client.R().
		SetHeader("Accept", "application/json").
		SetResult(&models.Schedule{}).
		Get("/api/v3/ride/live")
	return resp.Result().(*models.Schedule), nil
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

// Me method creates a request to retrieve data about the current user
func (c *ApiClient) Me() (*models.User, error) {
	resp, _ := c.Client.R().
		SetHeader("Accept", "application/json").
		SetResult(&models.User{}).
		Get("/api/me")
	return resp.Result().(*models.User), nil
}
