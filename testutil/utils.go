package testutil

import (
	"context"
	"database/sql"
	"gym-map/config"
	"gym-map/db"
	"math/rand"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
)

type FactoryOption[T any] func(*T)

func RandomInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000000)
}

func RandomUUID() string {
	return uuid.New().String()
}

func SetupDb(t *testing.T) *bun.Tx {
	config := config.GetConfig()
	db := db.GetDbConn(config.GetDSN(), true, "file://../migrations")
	db.DownMigrations()

	db.RunMigrations()
	tx, err := db.Conn.BeginTx(context.TODO(), &sql.TxOptions{})
	if err != nil {
		panic(err)
	}
	t.Cleanup(func() {
		tx.Rollback()
		db.DownMigrations()
		db.Conn.Close()
	})
	return &tx
}

func AssertCmpEqual[T any](t *testing.T, model T, actual, exepected any, opts ...cmp.Option) {
	t.Helper()
	options := append(opts, cmpopts.IgnoreFields(model, "Timestamp"))
	diff := cmp.Diff(actual, exepected, options...)
	if diff != "" {
		assert.Fail(t, "Values are not equal:\n"+diff)
	}

}
