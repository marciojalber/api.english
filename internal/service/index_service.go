// internal/routes/index_handler.go

package service

import (
	"fmt"
	"net/http"
)

func indexService(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		custom404(w, r.URL.Path)
		return
	}

	fmt.Fprintln(w, "Home")
}
