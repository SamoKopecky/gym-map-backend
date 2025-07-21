package schema

import "gym-map/model"

type Category struct {
	model.Category
	Properties []model.Property `json:"properties"`
}
