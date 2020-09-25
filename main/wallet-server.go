package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"wallet"
	"wallet/config"
	"wallet/registry"
	"wallet/rest"
	"wallet/storage/postgres"
)

func main() {
	// read yaml config file. Dont pass path to read
	// from default path
	cfg := config.ReadYaml("")
	confg := wallet.GetConfig(*cfg)

	database, err := postgres.NewDatabase(confg)
	if err != nil {
		log.Printf("database err %s", err)
		os.Exit(1)
	}

	// run migrations; update tables
	postgres.Migrate(database)

	channels := registry.NewChannels()
	domain := registry.NewDomain(confg, database, channels)

	// create the fiber server.
	server := fiber.New()
	rest.Router(server, domain, confg) // add endpoints

	// listen and serve
	port := fmt.Sprintf(":%v", 6700)
	log.Fatal(server.Listen(port))
}
