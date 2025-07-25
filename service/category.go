package service

import (
	"gym-map/model"
	"gym-map/store"
)

type Category struct {
	CategoryCrud store.Category
	PropertyCrud store.Property
	ExerciseCrud store.Exercise
}

func (c Category) GetByPropertyIds(properetyIds []int) (res []model.Category, err error) {
	properties, err := c.PropertyCrud.GetManyByIds(properetyIds)
	if err != nil {
		return
	}

	propertiesMap := make(map[int][]model.Property)
	for _, property := range properties {
		propertiesMap[property.CategoryId] = append(propertiesMap[property.CategoryId], property)
	}

	var categoryIds []int
	for key := range propertiesMap {
		categoryIds = append(categoryIds, key)
	}

	categories, err := c.CategoryCrud.GetManyByIds(categoryIds)
	if err != nil {
		return
	}

	res = make([]model.Category, len(categories))

	for i, category := range categories {
		if categoryProperties, ok := propertiesMap[category.Id]; ok {
			category.Properties = categoryProperties
		} else {
			category.Properties = []model.Property{}
		}
		res[i] = category
	}

	return
}
