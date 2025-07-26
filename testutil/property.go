package testutil

import (
	"gym-map/model"
	"testing"
)

func PropertyId(t *testing.T, id int) FactoryOption[model.Property] {
	t.Helper()
	return func(p *model.Property) {
		p.Id = id
	}
}

func PropertyCategoryId(t *testing.T, categoryId int) FactoryOption[model.Property] {
	t.Helper()
	return func(p *model.Property) {
		p.CategoryId = categoryId
	}
}

func PropertyFactory(t *testing.T, options ...FactoryOption[model.Property]) model.Property {
	t.Helper()

	property := model.Property{
		Name:       "Test Property",
		CategoryId: 1,
	}
	property.Id = RandomInt()

	for _, option := range options {
		option(&property)
	}
	return property
}

