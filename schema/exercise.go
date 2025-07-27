package schema

import "gym-map/model"

type Exercise struct {
	model.Exercise
	Categories       []model.Category `json:"categories"`
	InstructionCount int              `json:"instruction_count"`
}
