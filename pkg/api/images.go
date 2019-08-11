package api

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zachpuck/dokuweb-api/pkg/dropbox"
	"log"
	"net/http"
)

type ImageHandler struct {

}

func (h *ImageHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "GET":
		h.handleGet(res, req)
	default:
		http.Error(res, "Only GET are allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ImageHandler) handleGet(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "TODO: Get All Images from Dropbox")

	dpx := dropbox.New()
	images, err := dpx.ListFolder()
	if err != nil {
		http.Error(res, "Unable to get Images", http.StatusInternalServerError)
		log.Println(errors.Wrap(err, "unable to get images from dropbox"))
	}

	// convert images object to json string
	jsonBytes, err := json.Marshal(images)
	jsonString := string(jsonBytes)
	num, err := fmt.Fprintln(res, jsonString)
	if err != nil {
		log.Println(errors.Wrap(err, "something bad happened when converting images to json"))
	}
	fmt.Println("what is the num response for Fprintln? ", num)

	return
}