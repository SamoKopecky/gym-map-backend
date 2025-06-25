package store

import (
	"gym-map/model"
	"gym-map/schema"
)

type Machine interface {
	StoreBase[model.Machine]
	UpdatePosition(model *model.Machine) error
	GetWithCount() (machines []schema.Machine, err error)
}
