package dropbox

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"time"
)

func (d *DBX) ListFolder() ([]Entry, error) {
	// "https://api.dropboxapi.com/2/files/list_folder"

	jsonData := map[string]string{"path": ""}
	req, err := d.NewRequest(jsonData, "POST", "files/list_folder")
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	res, err := d.Client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	var result APIResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode response body")
	}

	fmt.Println("RESULTS: ", result.Entries[2].ID)

	return result.Entries, nil
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

func (d *DBX) GetThumbnail(id string) (GetThumbnailOutput, error) {
	//https://content.dropboxapi.com/2/files/get_thumbnail
	jsonData := map[string]string{
		"path": fmt.Sprintf("id:%s", id),
	}
	fmt.Println("Dat: ", jsonData)
	req, err := d.NewRequest(jsonData, "GET", "/files/get_thumbnail")
	if err != nil {
		return GetThumbnailOutput{}, errors.Wrap(err, "failed to get new request")
	}
	res, err := d.Client.Do(req)
	if err != nil {
		return GetThumbnailOutput{}, errors.Wrap(err, "error calling dropbox api")
	}
	var result GetThumbnailOutput
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return GetThumbnailOutput{}, errors.Wrap(err, "error decoding response body")
	}

	fmt.Println("GetThumbnail: ", result.Name)
	return result, nil
}

type GetThumbnailOutput struct {
	Name           string    `json:"name"`
	ID             string    `json:"id"`
	ClientModified time.Time `json:"client_modified"`
	ServerModified time.Time `json:"server_modified"`
	Rev            string    `json:"rev"`
	Size           int       `json:"size"`
	PathLower      string    `json:"path_lower"`
	PathDisplay    string    `json:"path_display"`
	SharingInfo    struct {
		ReadOnly             bool   `json:"read_only"`
		ParentSharedFolderID string `json:"parent_shared_folder_id"`
		ModifiedBy           string `json:"modified_by"`
	} `json:"sharing_info"`
	IsDownloadable bool `json:"is_downloadable"`
	PropertyGroups []struct {
		TemplateID string `json:"template_id"`
		Fields     []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"fields"`
	} `json:"property_groups"`
	HasExplicitSharedMembers bool   `json:"has_explicit_shared_members"`
	ContentHash              string `json:"content_hash"`
}
