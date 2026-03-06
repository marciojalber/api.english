package main

import (
	"fmt"
	"net/http"

	"github.com/marciojalber/api.english/internal/routes"
)

func main() {
	router 	:= routes.NewRouter()
	port 	:= 8080
	addr 	:= fmt.Sprintf(":%d", port)

	fmt.Printf("\nServer listening on http://localhost%s\n\n", addr)
	
	err 	:= http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
