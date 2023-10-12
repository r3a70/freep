package main

import (
	"log"
	"os"

	"freep.space/fsp/server"
)

const (
	ADDRESS string = ":8000"
)

func main() {

	if _, err := os.Stat("./downloads"); os.IsNotExist(err) {
		if err := os.Mkdir("downloads", 0755); err != nil {
			log.Println(err)
		}
	}
	server.Serve(ADDRESS)

}
