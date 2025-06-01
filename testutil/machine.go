package testutil

import (
	"gym-map/model"
	"testing"
)

func MachineFactory(t *testing.T, options ...FactoryOption[model.Machine]) model.Machine {
	t.Helper()
	machine := model.BuildMachine("name", "descrption", []string{"foo", "bar"}, 10, 10, 0, 0)
	machine.Id = RandomInt()

	for _, option := range options {
		option(&machine)
	}
	return machine
}
