package model

import "github.com/uptrace/bun"

type Category struct {
	bun.BaseModel `bun:"table:category"`
	IdModel
	Timestamp

	Name string `json:"name"`
}
