package routes

import (
    "fmt"
    "net/http"
    "time"
)

func NewRouter() http.Handler {
    mux := http.NewServeMux()

    mux.HandleFunc("/", indexHandler)
    mux.HandleFunc("/api/cards", apiCardsHandler)

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
        w.WriteHeader(http.StatusNotFound)
        next.ServeHTTP(w, r)
    })
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        custom404(w, r.URL.Path)
        return
    }

    fmt.Fprintln(w, "Index")
}

func custom404(w http.ResponseWriter, url string) {
    res := fmt.Sprintf(`{"err": "route_not_found", "txt": "The requested endpoint [%s] does not exist"}`, url)
    fmt.Fprint(w, res)
}
