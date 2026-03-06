package routes

import (
    "fmt"
    "net/http"
)

func NewRouter() http.Handler {
    mux := http.NewServeMux()

    mux.HandleFunc("/", indexHandler)
    mux.HandleFunc("/api/cards", apiCardsHandler)

    return logRequests(mux)
}

func logRequests(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        fmt.Printf(
            "%s -> %s %s\n",
            r.RemoteAddr,
            r.Method,
            r.URL.Path,
        )

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
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNotFound)

    res := fmt.Sprintf(`{"err": "route_not_found", "txt": "The requested endpoint [%s] does not exist"}`, url)
    fmt.Fprint(w, res)
}

func apiCardsHandler(w http.ResponseWriter, r *http.Request) {
    ctx := r.URL.Query().Get("context")
    var res string
    if ctx == "" {
        res = fmt.Sprintf("Cards of [%s] context", ctx)
    } else {
        res = "Context undefined for the cards"
    }
    
    fmt.Fprintln(w, res)
}
