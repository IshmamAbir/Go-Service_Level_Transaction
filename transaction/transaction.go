package transaction

import (
	"context"
	"database/sql"
	"log"

	"gorm.io/gorm"
)

type contextKey string

const txKey contextKey = "db_tx"

type UoW struct {
	db *gorm.DB
}

func NewUW(db *gorm.DB) UoW {
	return UoW{db}
}

func (uw *UoW) WithTx(ctx context.Context, fn func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	// begin a transaction
	tx := uw.db.Begin(&sql.TxOptions{})
	ctx = context.WithValue(ctx, txKey, tx)

	var done bool

	defer func() {
		if !done {
			tx.Rollback()
		}
	}()

	v, err := fn(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	done = true
	// Or commit the transaction
	tx.Commit()
	return v, nil
}

func GetTx(ctx context.Context) (tx *gorm.DB, ok bool) {
	tx, ok = ctx.Value(txKey).(*gorm.DB)
	return
}
