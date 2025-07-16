package store

import "gym-map/model"

type Media interface {
	StoreBase[model.Media]
	GetByIds(ids []int) ([]model.Media, error)
}
