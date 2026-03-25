// internal/handler/router.go

package router

import (
	"fmt"
	"net/http"
	"time"
	"strings"

	"github.com/marciojalber/api.english/internal/service/index_service"
	"github.com/marciojalber/api.english/internal/service/cards_service"
)

type statusWriter struct {
    http.ResponseWriter
    status int
    size   int
}

func (w *statusWriter) WriteHeader(code int) {
    w.status = code
    w.ResponseWriter.WriteHeader(code)
}

func (w *statusWriter) Write(b []byte) (int, error) {
    if w.status == 0 {
        w.status = http.StatusOK
    }
    n, err := w.ResponseWriter.Write(b)
    w.size += n
    return n, err
}

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index_service.Service)
	mux.HandleFunc("/api/cards", cards_service.Service)

	return logRequests(mux)
}

func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/favicon.ico" || strings.HasPrefix(r.URL.Path, "/.well-known/") {
		    w.WriteHeader(http.StatusNoContent)
		    return
		}

        start 	:= time.Now()
        reqID 	:= fmt.Sprintf("%d", start.UnixNano())
        sw 		:= &statusWriter{
            ResponseWriter: w,
            status:         0, // will default to 200 if not set
        }

		sw.Header().Set("Content-Type", "application/json")
		sw.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		sw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		sw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        // Handle CORS preflight properly
        if r.Method == http.MethodOptions {
            sw.WriteHeader(http.StatusNoContent)
            return
        }

		next.ServeHTTP(sw, r)

        duration := time.Since(start)

        // Ensure status is set (implicit 200 case)
        status := sw.status
        if status == 0 {
            status = http.StatusOK
        }

		fmt.Printf(
			"%s ... req_id=%s [%s / %db]\n   %s -> %d %s %s \n\n",
			start.Format("2006-01-02 15:04:05"),
			reqID,
			duration,
			sw.size,
			r.RemoteAddr,
			status,
			r.Method,
			r.URL.Path,
		)
	})
}
