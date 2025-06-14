package store

import "gym-map/model"

type Machine interface {
	StoreBase[model.Machine]
	UpdatePosition(model *model.Machine) error
}
