package crud

import (
	"gym-map/model"

	"github.com/uptrace/bun"
)

type Category struct {
	CRUDBase[model.Category]
}

func NewCategory(db bun.IDB) Category {
	return Category{CRUDBase: CRUDBase[model.Category]{db: db}}
}