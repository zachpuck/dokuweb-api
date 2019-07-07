package main

import (
	"fmt"
	"github.com/zachpuck/dokuweb-api/pkg/api"
)

func main() {
	fmt.Println("Starting Dokuforest web api")
	api.Start()
}
