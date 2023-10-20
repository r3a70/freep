package server

import (
	"fmt"
	"log"
	"net/http"

	"freep.space/fsp/internals"
	"freep.space/fsp/middlewares"
)

func redirectToTls(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://freep.space"+r.RequestURI, http.StatusMovedPermanently)
}

func Serve(address, addressHttp string, isOverTls bool) {

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

	go func() {
		if err := http.ListenAndServe(addressHttp, http.HandlerFunc(redirectToTls)); err != nil {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()

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
