package schema

import "gym-map/model"

type Exercise struct {
	model.Exercise
	InstructionCount int `json:"instruction_count"`
}
