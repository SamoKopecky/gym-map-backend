package store

import "gym-map/model"

type Exercise interface {
	StoreBase[model.Exercise]
	GetByMachineId(machineId int) (exercises []model.Exercise, err error)
	GetWithCount() (exercises []model.ExerciseWithCount, err error)
	GetWithCountMachineId(machineId int) (exercises []model.ExerciseWithCount, err error)
}
