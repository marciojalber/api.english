// internal/handler/router.go

package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/marciojalber/api.english/internal/service"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", service.IndexService)
	mux.HandleFunc("/api/cards", service.ApiCardsService)

	return logRequests(mux)
}

func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		now := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf(
			"%s ... %s -> %s %s\n",
			now,
			r.RemoteAddr,
			r.Method,
			r.URL.Path,
		)

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusNotFound)
		next.ServeHTTP(w, r)
	})
}
