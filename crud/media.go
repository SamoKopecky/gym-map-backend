package crud

import (
	"context"
	"gym-map/model"

	"github.com/uptrace/bun"
)

type Media struct {
	CRUDBase[model.Media]
}

func NewMedia(db bun.IDB) Media {
	return Media{CRUDBase: CRUDBase[model.Media]{db: db}}
}

func (m Media) GetByIds(ids []int) (medias []model.Media, err error) {
	err = m.db.NewSelect().
		Model(&medias).
		Where("id IN (?)", bun.In(ids)).
		Scan(context.Background())

	return
}
