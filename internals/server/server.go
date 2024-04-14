package server

import (
	"log"
	"net/http"

	"freep.space/fsp/internals/constant"
	"freep.space/fsp/internals/middlewares"
)

type Config struct {
	Address   string
	IsOverTls bool
	CertFile  string
	KeyFile   string
}

func (c Config) Run() {

	mux := http.NewServeMux()

	// Serving Static files
	homePage := http.FileServer(http.Dir("./static"))

	// Add handlers
	uploadHandler := http.HandlerFunc(Upload)
	downloadHandler := http.HandlerFunc(Download)
	ipHandler := http.HandlerFunc(Ip)

	// Registering Handlers
	mux.Handle("/", middlewares.Logger(middlewares.Security(homePage)))
	mux.Handle("/upload", middlewares.Logger(middlewares.Security(uploadHandler)))
	mux.Handle("/download/", middlewares.Logger(middlewares.Security(downloadHandler)))
	mux.Handle("/ip", middlewares.Logger(middlewares.Security(ipHandler)))

	// Show To user that the server is run properly
	log.Printf(constant.GREEN + "Freep Web server is running on " + c.Address + constant.RESET)

	var err any
	// listen and serve server at given address
	if c.IsOverTls {
		err = http.ListenAndServeTLS(c.Address, c.CertFile, c.KeyFile, mux)
	} else {
		err = http.ListenAndServe(c.Address, mux)
	}
	if err != nil {
		log.Fatalf(constant.RED+"There was a problem while running Freep server. the error is %v\n"+constant.RESET, err)
	}
}
