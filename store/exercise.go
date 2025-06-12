package store

import "gym-map/model"

type Exercise interface {
	StoreBase[model.Exercise]
}
