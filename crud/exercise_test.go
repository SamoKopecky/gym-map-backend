package crud

import (
	"gym-map/model"
	"gym-map/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetExerciseWithCount(t *testing.T) {
	db := testutil.SetupDb(t)
	crud := NewExercise(db)
	machineCrud := NewMachine(db)
	instructionCrud := NewInstruction(db)

	var expected []model.Exercise

	machine := testutil.MachineFactory(t)
	machineCrud.Insert(&machine)

	for range 2 {
		exercise := testutil.ExerciseFactory(t, testutil.ExerciseMachineId(t, machine.Id))
		if err := crud.Insert(&exercise); err != nil {
			t.Fatalf("Failed to insert exercise: %v", err)
		}
		expected = append(expected, exercise)
	}

	for range 3 {
		instruction := testutil.InstructionFactory(t, testutil.InstructionExerciseId(t, expected[0].Id))
		instructionCrud.Insert(&instruction)
	}

	// Act
	actual, err := crud.GetWithCountMachineId(machine.Id)
	if err != nil {
		t.Fatalf("Failed to get exercises: %v", err)
	}

	// Assert
	assert.Equal(t, len(expected), len(actual))
	assert.Equal(t, 3, actual[0].InstructionCount)
	assert.Equal(t, 0, actual[1].InstructionCount)
}

func TestGetExerciseWithCountWithMachineId(t *testing.T) {
	db := testutil.SetupDb(t)
	crud := NewExercise(db)
	machineCrud := NewMachine(db)
	instructionCrud := NewInstruction(db)

	var expected []model.Exercise
	var machines []model.Machine

	for range 2 {
		machine := testutil.MachineFactory(t)
		machineCrud.Insert(&machine)
		machines = append(machines, machine)
	}

	for i := range 2 {
		for range 2 {
			exercise := testutil.ExerciseFactory(t, testutil.ExerciseMachineId(t, machines[i].Id))
			if err := crud.Insert(&exercise); err != nil {
				t.Fatalf("Failed to insert exercise: %v", err)
			}
			expected = append(expected, exercise)

		}
	}

	for range 3 {
		instruction := testutil.InstructionFactory(t, testutil.InstructionExerciseId(t, expected[0].Id))
		instructionCrud.Insert(&instruction)
	}

	for range 1 {
		instruction := testutil.InstructionFactory(t, testutil.InstructionExerciseId(t, expected[3].Id))
		instructionCrud.Insert(&instruction)
	}

	// Act
	actual, err := crud.GetWithCountMachineId(machines[0].Id)
	if err != nil {
		t.Fatalf("Failed to get exercises: %v", err)
	}

	// Assert
	assert.Equal(t, 2, len(actual))
	assert.Equal(t, 3, actual[0].InstructionCount)
	assert.Equal(t, 0, actual[1].InstructionCount)
}
