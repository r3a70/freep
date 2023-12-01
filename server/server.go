package server

import (
	"fmt"
	"log"
	"net/http"

	"freep.space/fsp/internals"
	"freep.space/fsp/middlewares"
)

func Serve(address string, isOverTls bool) {

	mux := http.NewServeMux()

	// Serving Static files
	homePage := http.FileServer(http.Dir("./static"))

	// Add handlers
	uploadHandler := http.HandlerFunc(Upload)
	downloadHandler := http.HandlerFunc(Download)

	// Registering Handlers
	mux.Handle("/", middlewares.Logger(middlewares.Security(homePage)))
	mux.Handle("/upload", middlewares.Logger(middlewares.Security(uploadHandler)))
	mux.Handle("/download/", middlewares.Logger(middlewares.Security(downloadHandler)))

	// Show To user that the server is run properly
	fmt.Println(internals.GREEN + "Freep Web server is running on " + address + internals.RESET)

	var err any
	// listen and serve server at given address
	if isOverTls {
		err = http.ListenAndServeTLS(address, "freep.space.crt", "freep.space.key", mux)
	} else {
		err = http.ListenAndServe(address, mux)
	}
	if err != nil {
		log.Fatalf(internals.RED+"There was a problem while running Freep server. the error is %v\n"+internals.RESET, err)
	}
}
