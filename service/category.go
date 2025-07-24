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

func (c Category) GetCategoriesByExercise(exerciseId int) (categories []model.Category, err error) {
	exercise, err := c.ExerciseCrud.GetById(exerciseId)
	if err != nil || len(exercise.PropertyIds) == 0 {
		return
	}

	categories, err = c.CategoryCrud.GetCategoryProperties(&exercise.PropertyIds)
	if err != nil {
		return
	}

	return
}
