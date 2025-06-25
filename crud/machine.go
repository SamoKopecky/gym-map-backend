package crud

import (
	"context"
	"gym-map/model"
	"gym-map/schema"

	"github.com/uptrace/bun"
)

type Machine struct {
	CRUDBase[model.Machine]
}

func NewMachine(db bun.IDB) Machine {
	return Machine{CRUDBase: CRUDBase[model.Machine]{db: db}}
}

func (m Machine) UpdatePosition(model *model.Machine) error {
	// Don't omit zero
	query := m.db.NewUpdate().
		Model(model).
		Set("width = ?", model.Width).
		Set("height = ?", model.Height).
		Set("position_x = ?", model.PositionX).
		Set("position_y = ?", model.PositionY).
		WherePK()

	_, err := query.Exec(context.Background())
	return err
}

func (m Machine) GetWithCount() (machines []schema.Machine, err error) {
	subq := m.db.NewSelect().
		Model((*model.Machine)(nil)).
		ColumnExpr("machine.id AS machine_id").
		ColumnExpr("COUNT(exercise.id) as exercise_count").
		Join("JOIN exercise ON exercise.machine_id = machine.id").
		Group("machine.id")

	err = m.db.NewSelect().
		With("counts", subq).
		Model((*model.Machine)(nil)).
		ColumnExpr("machine.*").
		ColumnExpr("counts.exercise_count AS exercise_count").
		Join("LEFT JOIN counts on machine.id = counts.machine_id").
		Scan(context.Background(), &machines)

	return
}
