package crud

import (
	"gym-map/model"

	"github.com/uptrace/bun"
)

type Machine struct {
	CRUDBase[model.Machine]
}

func NewMachine(db bun.IDB) Machine {
	return Machine{CRUDBase: CRUDBase[model.Machine]{db: db}}
}
