package dropbox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	//"net/url"
	"github.com/pkg/errors"
	"os"
)

type DBX struct {
	Client *http.Client
	//BaseURL *url.URL
	BaseURL string
	Token   string
	Path    string
}

func New() *DBX {
	dbx := &DBX{}
	dbx.Client = &http.Client{}

	dbx.BaseURL = "https://api.dropboxapi.com/2"
	dbx.Token = os.Getenv("DROPBOX_APP_TOKEN")

	return dbx
}

func (d *DBX) NewRequest(data map[string]string, method string, apiPath string) (*http.Request, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode json")
	}

	req, err := http.NewRequest(
		method,
		fmt.Sprintf("%s/%s", d.BaseURL, apiPath),
		bytes.NewBuffer(b),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to complete request")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", d.Token))

	return req, nil
}
