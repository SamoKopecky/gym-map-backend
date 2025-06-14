package crud

import (
	"context"
	"gym-map/model"

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
