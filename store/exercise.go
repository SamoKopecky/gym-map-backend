package store

import "gym-map/model"

type Exercise interface {
	StoreBase[model.Exercise]
	GetByMachineId(machineId int) (exercises []model.Exercise, err error)
}
