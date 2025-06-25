package store

import (
	"gym-map/model"
	"gym-map/schema"
)

type Exercise interface {
	StoreBase[model.Exercise]
	GetWithCount() (exercises []schema.Exercise, err error)
	GetWithCountMachineId(machineId int) (exercises []schema.Exercise, err error)
}
