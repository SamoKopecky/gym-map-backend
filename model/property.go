package model

import "github.com/uptrace/bun"

type Property struct {
	bun.BaseModel `bun:"table:exercise"`
	IdModel
	Timestamp

	CategoryId int    `json:"category_id"`
	Name       string `json:"name"`
}
