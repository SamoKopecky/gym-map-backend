package store

import "gym-map/model"

type Category interface {
	StoreBase[model.Category]
	GetCategoryProperties() (categories []model.Category, err error)
}
