package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

func Start() {
	token := os.Getenv("DROPBOX_APP_TOKEN")
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "HomePage")
	})

	r.GET("/api/v1/all", func(c *gin.Context) {
		jsonData := map[string]string{"path": ""}
		jsonValue, _ := json.Marshal(jsonData)
		request, err := http.NewRequest("POST", "https://api.dropboxapi.com/2/files/list_folder", bytes.NewBuffer(jsonValue))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		client := &http.Client{}

		response, err := client.Do(request)
		if err != nil {
			fmt.Println("error getting all files: ", err)
		}
		data, _ := ioutil.ReadAll(response.Body)

		c.JSON(http.StatusOK, string(data))
	})

	r.Run()
}