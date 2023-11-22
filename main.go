package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// main Function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("<h1>This is the about page and 2 + 2 is %d.</h1>", sum))
}

// addValues adds two ints x and y, and returns the sum
func addValues(x, y int) int {
	return x + y
}
