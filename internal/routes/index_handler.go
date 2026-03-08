// internal/routes/index_handler.go

package routes

import (
    "net/http"
    "fmt"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        custom404(w, r.URL.Path)
        return
    }

    fmt.Fprintln(w, "Home")
}
