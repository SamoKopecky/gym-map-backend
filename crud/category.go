package crud

import (
	"context"
	"gym-map/model"

	"github.com/uptrace/bun"
)

type Category struct {
	CRUDBase[model.Category]
}

func NewCategory(db bun.IDB) Category {
	return Category{CRUDBase: CRUDBase[model.Category]{db: db}}
}

func (c Category) GetCategoryProperties() (categories []model.Category, err error) {
	err = c.db.NewSelect().
		Model(&categories).
		Relation("Properties").
		Scan(context.Background())

	return
}
