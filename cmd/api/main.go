// cmd/main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/marciojalber/api.english/internal/router"
	"github.com/marciojalber/api.english/internal/src"
)

func main() {
	cfg 	:= src.ConfigGet()
	router 	:= router.NewRouter()
	addr 	:= fmt.Sprintf(":%d", cfg.SERVER.Port)

	fmt.Printf("\nServer listening on http://localhost%s\n\n", addr)

	err 	:= http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
