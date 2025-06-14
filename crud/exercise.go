package crud

import (
	"context"
	"gym-map/model"

	"github.com/uptrace/bun"
)

type Exercise struct {
	CRUDBase[model.Exercise]
}

func NewExercise(db bun.IDB) Exercise {
	return Exercise{CRUDBase: CRUDBase[model.Exercise]{db: db}}
}

func (e Exercise) GetByMachineId(machineId int) (exercises []model.Exercise, err error) {
	err = e.db.NewSelect().
		Model(&exercises).
		Where("machine_id = ?", machineId).
		Scan(context.Background())

	return
}
