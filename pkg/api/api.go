package api

import (
	"fmt"
	"net/http"
	"github.com/zachpuck/dokuweb-api/pkg/util/urlpath"
)

type App struct {
	V1Handler *V1Handler


}

func (h *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = urlpath.ShiftPath(req.URL.Path)
	if head == "v1" {
		fmt.Println("v1 path: ", req.URL.Path)
		h.V1Handler.ServeHTTP(res, req)
		return
	}
	http.Error(res, "Not Found", http.StatusNotFound)
}