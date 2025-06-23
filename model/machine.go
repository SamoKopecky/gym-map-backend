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

// TODO: Add FKs and cascade on delete
// TOOD: Adjust get endpoints to use cruds that compute count everytime
// TODO: If the performance will be bad, split get into 2 endpoints
// /get -- just returns the tntieis
// /get/counts -- returuns only the counts, then combine them on FE
// USe counts in delete dialogs and card info bottom right corner probably or middle cause of decsription
