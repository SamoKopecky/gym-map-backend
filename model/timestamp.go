package model

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Timestamp struct {
	bun.BaseModel

	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"-"`
}

func (t *Timestamp) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		t.CreatedAt = time.Now().UTC()
		t.UpdatedAt = time.Now().UTC()
	case *bun.UpdateQuery:
		t.UpdatedAt = time.Now().UTC()
	}
	return nil
}
