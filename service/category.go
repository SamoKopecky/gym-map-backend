package service

import (
	"gym-map/model"
	"gym-map/schema"
	"gym-map/store"
)

type Category struct {
	CategoryCrud store.Category
	PropertyCrud store.Property
}

func (c Category) GetCategories() (schemaCategories []schema.Category, err error) {
	categories, err := c.CategoryCrud.Get()
	if err != nil {
		return
	}

	properties, err := c.PropertyCrud.Get()
	if err != nil {
		return
	}
	var propertiesMap map[int][]model.Property
	for _, property := range properties {
		if val, ok := propertiesMap[property.CategoryId]; ok {
			val = append(val, property)
		}
	}

	schemaCategories = make([]schema.Category, len(categories))

	for i, category := range categories {
		var schemaCategory schema.Category
		schemaCategory.Category = category
		if categoryProperties, ok := propertiesMap[category.Id]; ok {
			schemaCategory.Properties = categoryProperties
		} else {
			schemaCategory.Properties = []model.Property{}
		}
		schemaCategories[i] = schemaCategory
	}

	return
}
