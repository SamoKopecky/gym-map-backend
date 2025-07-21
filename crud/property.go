package crud

import (
	"gym-map/model"

	"github.com/uptrace/bun"
)

type Property struct {
	CRUDBase[model.Property]
}

func NewProperty(db bun.IDB) Property {
	return Property{CRUDBase: CRUDBase[model.Property]{db: db}}
}