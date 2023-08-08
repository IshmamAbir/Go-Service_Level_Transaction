package transaction

import (
	"context"
	"log"
	"net/http"

	"gorm.io/gorm"
)

func DBTransactionMiddleware(db *gorm.DB, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		txHandle := db.Begin()
		log.Print("beginning database transaction")

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "db_trx", txHandle)

		r = r.WithContext(ctx)

		// Create a custom ResponseWriter that captures the status
		customWriter := &responseWriter{ResponseWriter: w}
		handler(customWriter, r)

		status := customWriter.status
		if StatusInList(status, []int{http.StatusOK, http.StatusCreated}) {
			log.Print("committing transactions")
			if err := txHandle.Commit().Error; err != nil {
				log.Print("trx commit error: ", err)
			}
		} else {
			log.Print("rolling back transaction due to status code: ", status)
			txHandle.Rollback()
		}
	}
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
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
