package crud

import (
	"gym-map/model"

	"github.com/uptrace/bun"
)

type Exercise struct {
	CRUDBase[model.Exercise]
}

func NewExercise(db bun.IDB) Exercise {
	return Exercise{CRUDBase: CRUDBase[model.Exercise]{db: db}}
}
