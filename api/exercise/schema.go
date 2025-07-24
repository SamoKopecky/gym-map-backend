package exercise

import (
	"gym-map/api"
	"gym-map/model"
)

type exerciseGetRequest struct {
	MachineId *int `query:"machine_id"`
}

type exercisePostRequest struct {
	Name         string            `json:"name"`
	Description  *string           `json:"description"`
	MuscleGroups *[]string         `json:"muscle_groups"`
	MachineId    int               `json:"machine_id"`
	Difficulty   *model.Difficulty `json:"difficulty"`
	PropertyIds  []int             `json:"property_ids"`
}

func (epr exercisePostRequest) ToNewModel() model.Exercise {
	return model.BuildExercise(epr.Name, epr.Description, epr.MuscleGroups, epr.MachineId, epr.Difficulty, epr.PropertyIds)
}

type exercisePatchRequest struct {
	Name         *string           `json:"name"`
	Description  *string           `json:"description"`
	MuscleGroups *[]string         `json:"muscle_groups"`
	Difficulty   *model.Difficulty `json:"difficulty"`
	PropertyIds  []int             `json:"property_ids"`
}

func (epr exercisePatchRequest) ToExistingModel(id int) model.Exercise {
	return model.Exercise{
		IdModel:      model.IdModel{Id: id},
		Name:         api.DerefString(epr.Name),
		Description:  epr.Description,
		MuscleGroups: epr.MuscleGroups,
		Difficulty:   epr.Difficulty,
		PropertyIds:  epr.PropertyIds,
	}
}
