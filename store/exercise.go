package store

import "gym-map/model"

type Exercise interface {
	StoreBase[model.Exercise]
	GetWithCount() (exercises []model.ExerciseWithCount, err error)
	GetWithCountMachineId(machineId int) (exercises []model.ExerciseWithCount, err error)
}
