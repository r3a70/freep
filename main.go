package main

import (
	"freep.space/fsp/server"
)

const (
	ADDRESS string = ":8000"
)

func main() {

	server.Serve(ADDRESS)

}
