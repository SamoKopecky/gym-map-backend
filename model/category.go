package model

import "github.com/uptrace/bun"

type Category struct {
	bun.BaseModel `bun:"table:category"`
	IdModel
	Timestamp

	Name       string     `json:"name"`
	Properties []Property `bun:"rel:has-many,join:id=category_id" json:"properties"`
}
