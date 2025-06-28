package store

import (
	"gym-map/model"
)

type Instruction interface {
	StoreBase[model.Instruction]
	GetByExerciseId(exerciseId int) (instructions []model.Instruction, err error)
	GetByUserId(userId string) (instructions []model.Instruction, err error)
	SaveFile(id int, media_id int) error
}
