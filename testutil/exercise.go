package testutil

import (
	"gym-map/model"
	"testing"
)

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
