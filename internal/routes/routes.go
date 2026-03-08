package routes

import (
    "fmt"
    "net/http"
    "time"

    "github.com/marciojalber/api.english/pkg/utils"
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

func custom404(w http.ResponseWriter, url string) {
    res := utils.ToJson(utils.JsonMap{
        "err": "route_not_found",
        "txt": fmt.Sprintf("The requested endpoint [%s] does not exist", url),
    })
    // res := fmt.Sprintf(`{"err": "route_not_found", "txt": "The requested endpoint [%s] does not exist"}`, url)
    fmt.Fprint(w, res)
}
