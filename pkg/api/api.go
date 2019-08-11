package api

import (
	"fmt"
	"net/http"
	"github.com/zachpuck/dokuweb-api/pkg/util/urlpath"
)

type App struct {
	V1Handler *V1Handler
	//FileHandler *http.Handler

}

func (h *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = urlpath.ShiftPath(req.URL.Path)
	fmt.Fprintln(res,"HEAD=", head)

	if head == "v1" {
		h.V1Handler.ServeHTTP(res, req)
		return
	}
	http.Error(res, "Not Found", http.StatusNotFound)
}

type V1Handler struct {
}

func (h *V1Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Not Implemented!")
	return
}