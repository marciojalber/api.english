// internal/routes/index_handler.go

package service

import (
	"fmt"
	"net/http"

	"github.com/marciojalber/api.english/pkg/utils"
)

func IndexService(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		custom404(w, r.URL.Path)
		return
	}

	fmt.Fprintln(w, "Home")
}

func custom404(w http.ResponseWriter, url string) {
	res := utils.ToJson(utils.JsonMap{
		"err": "route_not_found",
		"txt": fmt.Sprintf("The requested endpoint [%s] does not exist", url),
	})
	// res := fmt.Sprintf(`{"err": "route_not_found", "txt": "The requested endpoint [%s] does not exist"}`, url)
	fmt.Fprint(w, res)
}
