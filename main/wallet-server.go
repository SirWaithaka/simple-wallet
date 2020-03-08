package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"

	"wallet/rest"
)

func main() {

	// create the fiber server.
	server := fiber.New()
	rest.Router(server) // add endpoints

	// listen and serve
	port := fmt.Sprintf(":%v", 6700)
	log.Fatal(server.Listen(port))
}
