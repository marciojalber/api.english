// cmd/main.go

package main

import (
    "fmt"
    "net/http"
    
    "github.com/marciojalber/api.english/internal/config"
    "github.com/marciojalber/api.english/internal/routes"
)

func main() {
    cfg     := config.Load()
    router  := routes.NewRouter()
    addr    := fmt.Sprintf(":%d", cfg.SERVER.Port)
    err     := http.ListenAndServe(addr, router)

    if err != nil {
        panic(err)
    }

    fmt.Printf("\nServer listening on http://localhost%s\n\n", addr)
}
