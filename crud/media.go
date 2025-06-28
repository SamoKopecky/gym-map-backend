package crud

import (
	"gym-map/model"

	"github.com/uptrace/bun"
)

type Media struct {
	CRUDBase[model.Media]
}

func NewMedia(db bun.IDB) Media {
	return Media{CRUDBase: CRUDBase[model.Media]{db: db}}
}
