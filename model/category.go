package model

import "github.com/uptrace/bun"

type Category struct {
	bun.BaseModel `bun:"table:exercise"`
	IdModel
	Timestamp

	Name string `json:"name"`
}
