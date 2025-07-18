package crud

import (
	"context"
	"gym-map/model"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
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

func (i Instruction) SaveMedia(id int, media_ids []int) error {
	_, err := i.db.NewUpdate().
		Model((*model.Instruction)(nil)).
		Set("media_ids = media_ids || ?", pgdialect.Array(media_ids)).
		Where("id = ?", id).
		Exec(context.Background())

	return err

}
