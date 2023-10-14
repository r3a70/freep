package main

import (
	"log"
	"os"

	"freep.space/fsp/server"
	"github.com/joho/godotenv"
)

const (
	ADDRESS string = ":443"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	if _, err := os.Stat("./downloads"); os.IsNotExist(err) {
		if err := os.Mkdir("downloads", 0755); err != nil {
			log.Println(err)
		}
	}
	server.Serve(ADDRESS)

}
