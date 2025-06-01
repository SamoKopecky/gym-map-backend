package crud

import (
	"context"
	"gym-map/model"
	"gym-map/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInsert(t *testing.T) {
	db := testutil.SetupDb(t)
	crud := NewMachine(db)
	expected := testutil.MachineFactory(t)

	// Act
	if err := crud.Insert(&expected); err != nil {
		t.Fatalf("Failed to insert machine: %v", err)
	}

	// Assert
	var actual []model.Machine
	if err := db.NewSelect().Model(&actual).Scan(context.TODO()); err != nil {
		t.Fatalf("Failed to retrieve machine: %v", err)
	}

	assert.Len(t, actual, 1)
	testutil.AssertCmpEqual(t, model.Machine{}, expected, actual[0])
}

func TestGet(t *testing.T) {
	db := testutil.SetupDb(t)
	crud := NewMachine(db)
	var expected []model.Machine
	for range 3 {
		machine := testutil.MachineFactory(t)
		if err := crud.Insert(&machine); err != nil {
			t.Fatalf("Failed to insert machine: %v", err)
		}
		expected = append(expected, machine)
	}

	// Act
	actual, err := crud.Get()
	if err != nil {
		t.Fatalf("Failed to get machines: %v", err)
	}

	// Assert
	assert.Equal(t, len(expected), len(actual))
	testutil.AssertCmpEqual(t, model.Machine{}, expected, actual)
}

func TestUpdate(t *testing.T) {
	db := testutil.SetupDb(t)
	crud := NewMachine(db)
	expected := testutil.MachineFactory(t)
	crud.Insert(&expected)

	// Act
	expected.Name = "foobar"
	if err := crud.Update(&expected); err != nil {
		t.Fatalf("Failed to update machine: %v", err)
	}

	// Assert
	var actual []model.Machine
	if err := db.NewSelect().Model(&actual).Scan(context.TODO()); err != nil {
		t.Fatalf("Failed to retrieve work sets: %v", err)
	}
	assert.Len(t, actual, 1)
	testutil.AssertCmpEqual(t, model.Machine{}, expected, actual[0])
}

func TestDelete(t *testing.T) {
	db := testutil.SetupDb(t)
	crud := NewMachine(db)
	var expected []model.Machine
	for range 2 {
		machine := testutil.MachineFactory(t)
		crud.Insert(&machine)
		expected = append(expected, machine)
	}

	// Act
	if err := crud.Delete(expected[0].Id); err != nil {
		t.Fatalf("Failed to delete machine: %v", err)
	}

	// Assert
	var actual []model.Machine
	if err := crud.db.NewSelect().Model(&actual).Scan(context.Background()); err != nil {
		t.Fatalf("Failed to get machine: %v", err)
	}
	require.Len(t, actual, 1, "number of db models is not correct")
	testutil.AssertCmpEqual(t, model.Machine{}, expected[1], actual[0])
}
