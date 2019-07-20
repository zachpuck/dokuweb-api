package dropbox

import (
	"fmt"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"os"
)

func NewClient() error {
	config := dropbox.Config{
		Token: os.Getenv("DROPBOX_APP_TOKEN"),

	}
	client := files.New(config)
	response, err := client.ListFolder(&files.ListFolderArg{
		Path: "",
	})
	if err != nil {
		return err
	}

	fmt.Println("response: ", response)
}