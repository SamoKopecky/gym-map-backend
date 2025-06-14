package machine

import (
	"gym-map/api"
	"gym-map/model"
)

type machinePatchRequest struct {
	Name         *string   `json:"name"`
	Description  *string   `json:"description"`
	MuscleGroups *[]string `json:"muscle_groups"`
}

type machinePositionsPatchRequest struct {
	Width     int `json:"width"`
	Height    int `json:"height"`
	PositionX int `json:"position_x"`
	PositionY int `json:"position_y"`
}

type machinePostRequest struct {
	Name         string    `json:"name"`
	Description  *string   `json:"description"`
	MuscleGroups *[]string `json:"muscle_groups"`
}

func (mpr machinePostRequest) ToNewModel() model.Machine {
	return model.BuildMachine(mpr.Name, mpr.Description, mpr.MuscleGroups, 150, 150, 0, 0)
}

func (mpr machinePatchRequest) ToExistingModel(id int) model.Machine {
	return model.Machine{
		IdModel:      model.IdModel{Id: id},
		Name:         api.DerefString(mpr.Name),
		Description:  mpr.Description,
		MuscleGroups: mpr.MuscleGroups,
	}
}

func (mppr machinePositionsPatchRequest) ToExistingModel(id int) model.Machine {
	return model.Machine{
		IdModel:   model.IdModel{Id: id},
		Width:     mppr.Width,
		Height:    mppr.Height,
		PositionX: mppr.PositionX,
		PositionY: mppr.PositionY,
	}
}
