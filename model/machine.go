package model

import "github.com/uptrace/bun"

type Machine struct {
	bun.BaseModel `bun:"table:machine"`
	IdModel
	Timestamp

	Name         string    `json:"name"`
	Description  *string   `json:"description"`
	MuscleGroups *[]string `json:"muscle_groups" bun:",array"`
	Width        int       `json:"width"`
	Height       int       `json:"height"`
	PositionX    int       `json:"position_x" bun:"position_x"`
	PositionY    int       `json:"position_y" bun:"position_y"`
}

type MachineWithCount struct {
	Machine
	ExerciseCount int `json:"exercise_count"`
}

func BuildMachine(name string, description *string, muscleGroups *[]string, width, height, positionX, positionY int) Machine {
	return Machine{
		Name:         name,
		Description:  description,
		MuscleGroups: muscleGroups,
		Width:        width,
		Height:       height,
		PositionX:    positionX,
		PositionY:    positionY,
	}
}
