package server

import (
	"fmt"
	"freep.space/fsp/internals"
	"freep.space/fsp/middlewares"
	"log"
	"net/http"
)

func Serve(address string) {

	// Serving Static files
	homePage := http.FileServer(http.Dir("./static"))

	// Registering Handlers
	http.Handle("/", middlewares.Logger(homePage))

	// Show To user that the server is run properly
	fmt.Println(internals.GREEN + "Freep Web server is running on :8000" + internals.RESET)

	// listen and serve server at given address
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatalf(internals.RED+"There was a problem while running Freep server. the error is %v\n"+internals.RESET, err)
	}
}
