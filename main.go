package main

import (
	"log"

	"github.com/wcharczuk/instabot/server"
)

func main() {
	log.Fatal(server.Init().Start())
}
