package main

import (
	"fmt"
	"github.com/zachpuck/dokuweb-api/pkg/api"
	"log"
	"net/http"
)

const (
	servicePort = 8080
)

func main() {
	fmt.Printf("Starting Dokuforest web api on port: %d\n", servicePort)

	a := &api.App{
		V1Handler: new(api.V1Handler),
	}

	err := http.ListenAndServe(fmt.Sprintf(":%d", servicePort), a)
	if err != nil {
		log.Fatal(err)
	}
}
