package service

import (
	"gym-map/model"
	"gym-map/schema"
	store "gym-map/store/mock"
	"gym-map/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetByExercises(t *testing.T) {
	categoryCrud := store.NewMockCategory(t)
	propertyCrud := store.NewMockProperty(t)
	service := Category{
		CategoryCrud: categoryCrud,
		PropertyCrud: propertyCrud,
	}

	// Arrange
	var exercises []schema.Exercise
	// Single category, multiple properties
	// Mutliple categories
	// No categories & properties
	propertyIds := [][]int{{1, 2}, {3, 4}, {}}
	for i := range 3 {
		exercises = append(exercises, schema.Exercise{
			Exercise: testutil.ExerciseFactory(t,
				testutil.ExerciseId(t, (i+1)*100),
				testutil.ExercisePropertyIds(t, propertyIds[i])),
		})
	}
	var properties []model.Property
	categoryIds := []int{1, 1, 1, 2}
	for i := range 4 {
		properties = append(properties, testutil.PropertyFactory(t,
			testutil.PropertyId(t, i+1),
			testutil.PropertyCategoryId(t, categoryIds[i])))
	}

	var categories []model.Category
	for i := range 2 {
		categories = append(categories, testutil.CategoryFactory(t,
			testutil.CategoryId(t, i+1)))
	}

	propertyCrud.EXPECT().GetManyByIds(mock.Anything).Return(properties, nil)
	categoryCrud.EXPECT().GetManyByIds(mock.Anything).Return(categories, nil)

	// Act
	res, err := service.GetByExercises(exercises)

	// Assert
	require.Equal(t, err, nil)

	cat_100, ok := res[100]
	require.True(t, ok)
	assert.Len(t, cat_100, 1)
	assert.Equal(t, cat_100[0].Id, 1)
	assert.Len(t, cat_100[0].Properties, 2)
	assert.Equal(t, cat_100[0].Properties[0].Id, 1)
	assert.Equal(t, cat_100[0].Properties[0].CategoryId, 1)
	assert.Equal(t, cat_100[0].Properties[1].Id, 2)
	assert.Equal(t, cat_100[0].Properties[1].CategoryId, 1)

	cat_200, ok := res[200]
	require.True(t, ok)
	assert.Len(t, cat_200, 2)
	assert.Equal(t, cat_200[0].Id, 1)
	assert.Equal(t, cat_200[1].Id, 2)
	assert.Len(t, cat_200[0].Properties, 1)
	assert.Equal(t, cat_200[0].Properties[0].Id, 3)
	assert.Equal(t, cat_200[0].Properties[0].CategoryId, 1)
	assert.Len(t, cat_200[1].Properties, 1)
	assert.Equal(t, cat_200[1].Properties[0].Id, 4)
	assert.Equal(t, cat_200[1].Properties[0].CategoryId, 2)
}
