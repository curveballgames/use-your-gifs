// Package main is the entry point for the application
package main

import (
	"fmt"
	"net/http"

	"github.com/curveballgames/use-your-gifs/handlers"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/create", handlers.HandleNewRoom)

	appengine.Main()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprint(w, "Hello, World!")
}
