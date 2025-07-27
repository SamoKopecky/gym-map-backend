package testutil

import (
	"gym-map/model"
	"testing"
)

func CategoryId(t *testing.T, id int) FactoryOption[model.Category] {
	t.Helper()
	return func(p *model.Category) {
		p.Id = id
	}
}

func CategoryFactory(t *testing.T, options ...FactoryOption[model.Category]) model.Category {
	t.Helper()

	category := model.Category{
		Name: "Test Category",
	}
	category.Id = RandomInt()

	for _, option := range options {
		option(&category)
	}
	return category
}

