package testutil

import (
	"gym-map/model"
	"testing"
)

func ExercisePropertyIds(t *testing.T, propertyids []int) FactoryOption[model.Exercise] {
	t.Helper()
	return func(e *model.Exercise) {
		e.PropertyIds = propertyids
	}
}

func ExerciseId(t *testing.T, id int) FactoryOption[model.Exercise] {
	t.Helper()
	return func(e *model.Exercise) {
		e.Id = id
	}
}

func ExerciseMachineId(t *testing.T, machineId int) FactoryOption[model.Exercise] {
	t.Helper()
	return func(e *model.Exercise) {
		e.MachineId = machineId
	}
}

func ExerciseFactory(t *testing.T, options ...FactoryOption[model.Exercise]) model.Exercise {
	t.Helper()
	description := "foobar"

	exercise := model.BuildExercise("name", &description, 10, nil, []int{})
	exercise.Id = RandomInt()

	for _, option := range options {
		option(&exercise)
	}
	return exercise
}
