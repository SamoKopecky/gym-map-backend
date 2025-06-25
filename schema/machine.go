package schema

import "gym-map/model"

type Machine struct {
	model.Machine
	ExerciseCount int `json:"exercise_count"`
}
