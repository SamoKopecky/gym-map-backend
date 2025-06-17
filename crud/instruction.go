package crud

import (
	"context"
	"gym-map/model"

	"github.com/uptrace/bun"
)

type Instruction struct {
	CRUDBase[model.Instruction]
}

func NewInstruction(db bun.IDB) Instruction {
	return Instruction{CRUDBase: CRUDBase[model.Instruction]{db: db}}
}

func (i Instruction) GetByExerciseId(exerciseId int) (instructions []model.Instruction, err error) {
	err = i.db.NewSelect().
		Model(&instructions).
		Where("exercise_id = ?", exerciseId).
		Scan(context.Background())

	return
}

func (i Instruction) GetByUserId(userId string) (instructions []model.Instruction, err error) {
	err = i.db.NewSelect().
		Model(&instructions).
		Where("user_id = ?", userId).
		Scan(context.Background())

	return
}

func (i Instruction) CreateFile(id int, fileId, fileName string) error {
	_, err := i.db.NewUpdate().
		Model((*model.Instruction)(nil)).
		Set("file_id = ?", fileId).
		Set("file_name = ?", fileName).
		Where("id = ?", id).
		Exec(context.Background())

	return err

}
