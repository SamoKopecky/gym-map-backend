package machine

import (
	"gym-map/model"
)

type machinePostRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func (mpr machinePostRequest) ToModel() model.Machine {
	return model.BuildMachine(mpr.Name, mpr.Description, []string{}, 0, 0, 0, 0)
}
