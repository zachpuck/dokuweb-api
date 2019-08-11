package api

import (
	"fmt"
	"github.com/zachpuck/dokuweb-api/pkg/util/urlpath"
	"net/http"
)

type V1Handler struct {
	ImageHandler *ImageHandler
}

func (h *V1Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	if req.URL.Path != "/" {
		head, req.URL.Path = urlpath.ShiftPath(req.URL.Path)
		switch head {
		case "images":
			fmt.Println("images path: ", req.URL.Path)
			h.ImageHandler.ServeHTTP(res, req)
		default:
			http.Error(res, "Not Found", http.StatusNotFound)
		}
		return
	}
	//
	switch req.Method {
	case "GET":
		h.handleGet(res, req)
	default:
		http.Error(res, "Only GET are allowed", http.StatusMethodNotAllowed)
	}
}

func (h *V1Handler) handleGet(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "v1 root")
	return
}

