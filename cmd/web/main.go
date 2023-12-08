package main

import (
	"fmt"
	"net/http"

	handlers "github.com/kunovon/go-kgv-adminstration/pkg/handler"
)

const portNumber = ":8080"

// main is the main function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/impressum", handlers.Impressum)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}