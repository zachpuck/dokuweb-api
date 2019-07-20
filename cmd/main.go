package main

import (
	"fmt"
	"github.com/zachpuck/dokuweb-api/pkg/dropbox"
)

func main() {
	fmt.Println("Starting Dokuforest web api")
	dropbox.Start()
}
