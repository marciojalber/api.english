// internal/routes/index_service.go

package service

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func Service(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		custom404(w, r.URL.Path)
		return
	}

	fmt.Fprintln(w, "Home")
}

func custom404(w http.ResponseWriter, url string) {
	res, _ := json.Marshal(map[string]string{
		"err": "route_not_found",
		"txt": fmt.Sprintf("The requested endpoint [%s] does not exist", url),
	})
	fmt.Fprint(w, string(res))
}
