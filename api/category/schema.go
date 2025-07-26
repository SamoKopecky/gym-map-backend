package category

import (
	"gym-map/model"
)

type categoryPostRequest struct {
	Name string `json:"name"`
}

func (cpr categoryPostRequest) ToNewModel() model.Category {
	return model.Category{
		Name: cpr.Name,
	}
}

type categoryPatchRequest struct {
	Name string `json:"name"`
}

func (cpr categoryPatchRequest) ToExistingModel(id int) model.Category {
	return model.Category{
		IdModel: model.IdModel{Id: id},
		Name:    cpr.Name,
	}
}
