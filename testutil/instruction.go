package testutil

import (
	"gym-map/model"
	"testing"
)

func InstructionExerciseId(t *testing.T, exerciseId int) FactoryOption[model.Instruction] {
	t.Helper()
	return func(e *model.Instruction) {
		e.ExerciseId = exerciseId
	}
}

func InstructionFactory(t *testing.T, options ...FactoryOption[model.Instruction]) model.Instruction {
	t.Helper()

	instruction := model.BuildInstruction("123", "foobar", 0, nil, nil)
	instruction.Id = RandomInt()

	for _, option := range options {
		option(&instruction)
	}
	return instruction
}
