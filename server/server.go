package server

import (
	"fmt"
	"log"
	"net/http"
)

func Serve(address string) {

	// Serving Static files
	homePage := http.FileServer(http.Dir("./static"))

	// Registering Handlers
	http.Handle("/", homePage)

	// Show To user that the server is run properly
	fmt.Println("Starting freep server at :8000")

	// listen and serve server at given address
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatalf("There was a problem while running Freep server. the error is %v\n", err)
	}
}
