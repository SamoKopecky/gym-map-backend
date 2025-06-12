package model

import "github.com/uptrace/bun"

type Exercise struct {
	bun.BaseModel `bun:"table:exercise"`
	IdModel
	Timestamp

	MachineId    int       `json:"machine_id"`
	Name         string    `json:"name"`
	Description  *string   `json:"description"`
	MuscleGroups *[]string `json:"muscle_groups" bun:",array"`
}

func BuildExercise(name string, description *string, muscleGroups *[]string, machineId int) Exercise {
	return Exercise{
		Name:         name,
		Description:  description,
		MuscleGroups: muscleGroups,
		MachineId:    machineId,
	}
}
