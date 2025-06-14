package testutil

import (
	"gym-map/model"
	"testing"
)

func MachineFactory(t *testing.T, options ...FactoryOption[model.Machine]) model.Machine {
	t.Helper()
	description := "foobar"
	muscleGroups := []string{"foo", "bar"}

	machine := model.BuildMachine("name", &description, &muscleGroups, 10, 10, 0, 0)
	machine.Id = RandomInt()

	for _, option := range options {
		option(&machine)
	}
	return machine
}
