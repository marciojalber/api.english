//go:generate go run ../generator/main.go
// cmd/main.go

package main

import (
    "fmt"
    "net/http"
    
    "github.com/marciojalber/api.english/internal/src"
    "github.com/marciojalber/api.english/internal/router"
)

func main() {
    cfg     := src.Config.Load()
    router  := router.NewRouter()
    addr    := fmt.Sprintf(":%d", cfg.SERVER.Port)
    err     := http.ListenAndServe(addr, router)

    if err != nil {
        panic(err)
    }

    fmt.Printf("\nServer listening on http://localhost%s\n\n", addr)
}
