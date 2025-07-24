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

func (c Category) GetCategoryProperties(propertyIds *[]int) (categories []model.Category, err error) {
	query := c.db.NewSelect().
		Model(&categories).
		Relation("Properties")

	if propertyIds != nil {
		query = query.Where("property.id in (?)", bun.In(propertyIds))
	}
	err = query.Scan(context.Background())
	return
}
