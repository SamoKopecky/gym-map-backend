package store

import "gym-map/model"

type Category interface {
	StoreBase[model.Category]
	GetCategoryProperties(propertyIds *[]int) ([]model.Category, error)
}

