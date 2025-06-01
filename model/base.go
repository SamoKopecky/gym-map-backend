package model

import "github.com/uptrace/bun"

const EMPTY_ID = 0

type IdModel struct {
	bun.BaseModel

	Id int `bun:",pk,autoincrement" json:"id"`
}

func (im IdModel) IsEmpty() bool {
	return im.Id == EMPTY_ID
}
