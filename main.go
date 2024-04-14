package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"freep.space/fsp/internals/server"
	"github.com/joho/godotenv"
)

func main() {

	var config server.Config

	flag.StringVar(&config.Address, "addr", ":8000", "Address and Port the freep server should listen on")
	flag.BoolVar(&config.IsOverTls, "overtls", false, "is freep server runs over tls")
	flag.StringVar(&config.CertFile, "cert", "", "Path to the cert file")
	flag.StringVar(&config.KeyFile, "key", "", "Path to the key file")

	flag.Parse()

	if !strings.Contains(config.Address, ":") {
		config.Address = ":" + config.Address
	}
	if config.IsOverTls && (config.KeyFile == "" || config.CertFile == "") {
		log.Fatal("You must specify key and cert paths")
	}

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	if _, err := os.Stat("./downloads"); os.IsNotExist(err) {
		if err := os.Mkdir("downloads", 0755); err != nil {
			log.Println(err)
		}
	}
	config.Run()

}
