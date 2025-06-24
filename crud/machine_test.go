package crud

import (
	"gym-map/model"
	"gym-map/testutil"
	"gym-map/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMachineWithCount(t *testing.T) {
	db := testutil.SetupDb(t)
	exerciseCrud := NewExercise(db)
	machineCrud := NewMachine(db)

	var expected []model.Machine

	for range 2 {
		machine := testutil.MachineFactory(t)
		machineCrud.Insert(&machine)
		expected = append(expected, machine)
	}

	for i := range 2 {
		for range 2 + i {
			exercise := testutil.ExerciseFactory(t, testutil.ExerciseMachineId(t, expected[i].Id))
			if err := exerciseCrud.Insert(&exercise); err != nil {
				t.Fatalf("Failed to insert exercise: %v", err)
			}
		}

	}

	// Act
	actual, err := machineCrud.GetWithCount()
	if err != nil {
		t.Fatalf("Failed to get exercises: %v", err)
	}
	utils.PrettyPrint(actual)

	// Assert
	assert.Equal(t, len(expected), len(actual))
	assert.Equal(t, 2, actual[0].ExerciseCount)
	assert.Equal(t, expected[0].Id, actual[0].Id)
	assert.Equal(t, 3, actual[1].ExerciseCount)
	assert.Equal(t, expected[1].Id, actual[1].Id)
}
