package property

import (
	"gym-map/model"
)

type propertyPostRequest struct {
	CategoryId int    `json:"category_id"`
	Name       string `json:"name"`
}

func (ppr propertyPostRequest) ToNewModel() model.Property {
	return model.Property{
		CategoryId: ppr.CategoryId,
		Name:       ppr.Name,
	}
}

type propertyPatchRequest struct {
	Name string `json:"name"`
}

func (cpr propertyPatchRequest) ToExistingModel(id int) model.Property {
	return model.Property{
		IdModel: model.IdModel{Id: id},
		Name:    cpr.Name,
	}
}
