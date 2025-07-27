package store

import "gym-map/model"

type Property interface {
	StoreBase[model.Property]
}