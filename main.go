package main

import (
	"fmt"
	"net/http"

	"github.com/kixsian/gobox/pkg/handlers"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Root)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting Application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
