package service

import (
	mapset "github.com/deckarep/golang-set/v2"
	"gym-map/model"
	"gym-map/schema"
	"gym-map/store"
)

type Category struct {
	CategoryCrud store.Category
	PropertyCrud store.Property
}

func (c Category) GetByExercises(exercises []schema.Exercise) (res map[int][]model.Category, err error) {
	res = make(map[int][]model.Category)
	uniquePropertyIds := mapset.NewSet[int]()
	for _, exercise := range exercises {
		for _, propertyId := range exercise.PropertyIds {
			uniquePropertyIds.Add(propertyId)
		}
	}

	properties, err := c.PropertyCrud.GetManyByIds(uniquePropertyIds.ToSlice())
	if err != nil {
		return
	}
	categoryIds := mapset.NewSet[int]()
	propertiesMap := make(map[int]model.Property)
	for _, property := range properties {
		categoryIds.Add(property.CategoryId)
		propertiesMap[property.Id] = property
	}

	categories, err := c.CategoryCrud.GetManyByIds(categoryIds.ToSlice())
	if err != nil {
		return
	}

	categoriesMap := make(map[int]model.Category)
	for _, category := range categories {
		categoriesMap[category.Id] = category
	}

	for _, exercise := range exercises {
		// First group properties by category id
		iCategoryMap := make(map[int][]model.Property)
		for _, propertyId := range exercise.PropertyIds {
			if property, ok := propertiesMap[propertyId]; ok {
				iCategoryMap[property.CategoryId] = append(iCategoryMap[property.CategoryId], property)
			}
		}

		// Get unique categories and fill with properties
		for categoryId, properties := range iCategoryMap {
			if newCategory, ok := categoriesMap[categoryId]; ok {
				newCategory.Properties = properties
				res[exercise.Id] = append(res[exercise.Id], newCategory)
			}
		}
	}

	return
}
