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

func (e Exercise) GetWithCount() (exercises []model.ExerciseWithCount, err error) {
	query := e.getWithCountQuery()
	err = query.Scan(context.Background(), &exercises)
	return
}

func (e Exercise) GetWithCountMachineId(machineId int) (exercises []model.ExerciseWithCount, err error) {
	query := e.getWithCountQuery().Where("exercise.machine_id = ?", machineId)
	err = query.Scan(context.Background(), &exercises)
	return
}

func (e Exercise) getWithCountQuery() *bun.SelectQuery {
	subq := e.db.NewSelect().
		Model((*model.Exercise)(nil)).
		ColumnExpr("exercise.id AS exercise_id").
		ColumnExpr("COUNT(instruction.id) as instruction_count").
		Join("JOIN instruction ON instruction.exercise_id = exercise.id").
		Group("exercise.id")

	return e.db.NewSelect().
		With("counts", subq).
		Model((*model.Exercise)(nil)).
		ColumnExpr("exercise.*").
		ColumnExpr("counts.instruction_count AS instruction_count").
		Join("LEFT JOIN counts on exercise.id = counts.exercise_id")
}

func (e Exercise) GetByMachineId(machineId int) (exercises []model.Exercise, err error) {
	err = e.db.NewSelect().
		Model(&exercises).
		Where("machine_id = ?", machineId).
		Scan(context.Background())

	return
}
