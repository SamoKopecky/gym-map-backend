package model

import "github.com/uptrace/bun"

type Difficulty string

const (
	Easy   Difficulty = "easy"
	Normal            = "normal"
	Hard              = "hard"
	Empty             = ""
)

type Exercise struct {
	bun.BaseModel `bun:"table:exercise"`
	IdModel
	Timestamp

	MachineId    int         `json:"machine_id"`
	Name         string      `json:"name"`
	Description  *string     `json:"description"`
	MuscleGroups *[]string   `json:"muscle_groups" bun:",array"`
	Difficulty   *Difficulty `json:"difficulty"`
	PropertyIds  []int       `json:"property_ids" bun:",array"`
}

func BuildExercise(name string, description *string, muscleGroups *[]string, machineId int, difficulty *Difficulty, propertyIds []int) Exercise {
	return Exercise{
		Name:         name,
		Description:  description,
		MuscleGroups: muscleGroups,
		MachineId:    machineId,
		Difficulty:   difficulty,
		PropertyIds:  propertyIds,
	}
}
