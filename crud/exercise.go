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

const getWithCountQuery = `
WITH instruction_counts AS (
	 select count(instruction.id) as c, instruction.exercise_id from exercise join instruction on instruction.exercise_id = exercise.id group by instruction.exercise_id
) 
SELECT exercise.*, instruction_counts.c FROM exercise LEFT JOIN instruction_counts on exercise.id = instruction_counts.exercise_id
`

func (e Exercise) GetWithCount() (exercises []model.Exercise, err error) {
	err = e.db.NewRaw()

}

func (e Exercise) GetByMachineId(machineId int) (exercises []model.Exercise, err error) {
	err = e.db.NewSelect().
		Model(&exercises).
		Where("machine_id = ?", machineId).
		Scan(context.Background())

	return
}
