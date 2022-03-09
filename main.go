package main

import (
	//"fmt"
	"log"
	"github.com/Elbi123/learnmicro/server"
)

func main() {
	var endpoint = ":8085"
	//fmt.Printf("Server started on: %s", endpoint)
	log.Fatal(server.ServeAPI(endpoint))
}
