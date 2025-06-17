package store

import (
	"gym-map/model"
)

type Instruction interface {
	StoreBase[model.Instruction]
	GetByExerciseId(exerciseId int) (instructions []model.Instruction, err error)
	GetByUserId(userId string) (instructions []model.Instruction, err error)
	CreateFile(id int, fileId, fileName string) error
}
