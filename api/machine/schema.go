package machine

import (
	"gym-map/api"
	"gym-map/model"
)

type machinePatchRequest struct {
	Name         *string   `json:"name"`
	Description  *string   `json:"description"`
	MuscleGroups *[]string `json:"muscle_groups"`
	Width        *int      `json:"width"`
	Height       *int      `json:"height"`
	PositionX    *int      `json:"position_x"`
	PositionY    *int      `json:"position_y"`
}

type machinePostRequest struct {
	Name         string    `json:"name"`
	Description  *string   `json:"description"`
	MuscleGroups *[]string `json:"muscle_groups"`
}

func (mpr machinePostRequest) ToNewModel() model.Machine {
	return model.BuildMachine(mpr.Name, mpr.Description, mpr.MuscleGroups, 0, 0, 0, 0)
}

func (mpr machinePatchRequest) ToExistingModel(id int) model.Machine {
	return model.Machine{
		IdModel:      model.IdModel{Id: id},
		Name:         api.DerefString(mpr.Name),
		Description:  mpr.Description,
		MuscleGroups: mpr.MuscleGroups,
		Width:        api.DerefInt(mpr.Width),
		Height:       api.DerefInt(mpr.Height),
		PositionX:    api.DerefInt(mpr.PositionX),
		PositionY:    api.DerefInt(mpr.PositionY),
	}
}
