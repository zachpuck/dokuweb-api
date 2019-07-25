package dropbox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	//"net/url"
	"github.com/pkg/errors"
	"os"
	"time"
)

type DBX struct {
	Client *http.Client
	//BaseURL *url.URL
	BaseURL string
	Token   string
	Path    string
}

type APIResponse struct {
	Entries []Entry `json:"entries"`
	Cursor  string  `json:"cursor"`
	HasMore bool    `json:"has_more"`
}

type Entry struct {
	Tag            string    `json:".tag"`
	Name           string    `json:"name"`
	PathLower      string    `json:"path_lower"`
	PathDisplay    string    `json:"path_display"`
	ID             string    `json:"id"`
	ClientModified time.Time `json:"client_modified,omitempty"`
	ServerModified time.Time `json:"server_modified,omitempty"`
	Rev            string    `json:"rev,omitempty"`
	Size           int       `json:"size,omitempty"`
	IsDownloadable bool      `json:"is_downloadable,omitempty"`
	ContentHash    string    `json:"content_hash,omitempty"`
}

func New() *DBX {
	dbx := &DBX{}
	dbx.Client = &http.Client{}

	dbx.BaseURL = "https://api.dropboxapi.com/2"
	dbx.Token = os.Getenv("DROPBOX_APP_TOKEN")

	dbx.Path = ""

	return dbx
}

func (d *DBX) ListFolder() {
	// "https://api.dropboxapi.com/2/files/list_folder"

	jsonData := map[string]string{"path": d.Path}
	results, err := d.NewRequest(jsonData, "POST", "files/list_folder")
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	//jsonBytes, err := json.Marshal(jsonData)
	//url := fmt.Sprintf("%s/%s", d.BaseURL, "files/list_folder")
	//fmt.Println(url)
	//request, _ := http.NewRequest(
	//	"POST",
	//	url,
	//	bytes.NewBuffer(jsonBytes),
	//)
	//request.Header.Set("Content-Type", "application/json")
	//request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", d.Token))
	//
	//resp, err := d.Client.Do(request)
	//if err != nil {
	//	fmt.Println("error getting all files: ", err)
	//}
	//
	//var results APIResponse
	//decodeError := json.NewDecoder(resp.Body).Decode(&results)
	//if decodeError != nil {
	//	fmt.Println("ERR DECODING: ", decodeError)
	//}

	fmt.Println("RESULTS: ", results.Entries[0].Name)
}

func (d *DBX) NewRequest(data map[string]string, method string, apiPath string) (APIResponse, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return APIResponse{}, errors.Wrap(err, "failed to encode json")
	}

	req, err := http.NewRequest(
		method,
		fmt.Sprintf("%s/%s", d.BaseURL, apiPath),
		bytes.NewBuffer(b),
	)
	if err != nil {
		return APIResponse{}, errors.Wrap(err, "failed to complete request")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", d.Token))

	res, err := d.Client.Do(req)
	if err != nil {
		return APIResponse{}, errors.Wrap(err, "request failed")
	}

	var result APIResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return APIResponse{}, errors.Wrap(err, "failed to decode response body")
	}

	return result, nil
}
