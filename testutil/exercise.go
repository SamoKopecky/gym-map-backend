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
	muscleGroups := []string{"foo", "bar"}

	exercise := model.BuildExercise("name", &description, &muscleGroups, 10, nil)
	exercise.Id = RandomInt()

	for _, option := range options {
		option(&exercise)
	}
	return exercise
}
