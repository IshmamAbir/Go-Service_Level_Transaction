package transaction

import (
	"context"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type contextKey string

const txKey contextKey = "db_tx"

func DBTransactionMiddleware(db *gorm.DB, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		txHandle := db.Begin()
		log.Print("beginning database transaction")

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
				log.Panic(r)
			}
		}()

		ctx := SetTxKey(r.Context(), txHandle)
		r = r.WithContext(ctx)

		// Create a custom ResponseWriter that captures the status
		customWriter := &ResponseWriter{
			ResponseWriter: w,
		}
		handler(customWriter, r)

		status := customWriter.Status
		if StatusInList(status, []int{http.StatusOK, http.StatusCreated}) {
			log.Print("committing transactions")
			if err := txHandle.Commit().Error; err != nil {
				log.Print("trx commit error: ", err)
			}
		} else {
			log.Print("rolling back transaction due to an error ")
			txHandle.Rollback()
		}
	}
}

func SetTxKey(ctx context.Context, txHandle *gorm.DB) context.Context {
	return context.WithValue(ctx, txKey, txHandle)
}

func GetTxKey(ctx context.Context) *gorm.DB {
	return ctx.Value(txKey).(*gorm.DB)
}

type ResponseWriter struct {
	http.ResponseWriter
	Status int
}

func (rw *ResponseWriter) WriteHeader(code int) {
	rw.Status = code
	rw.ResponseWriter.WriteHeader(code)
}

func StatusInList(status int, statusList []int) bool {
	for _, s := range statusList {
		if status == s {
			return true
		}
	}
	return false
}
