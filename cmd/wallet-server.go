package main

import (
	"fmt"
	"log"
	"os"

	"simple-wallet/app"
	"simple-wallet/app/registry"
	"simple-wallet/app/rest"
	"simple-wallet/app/storage/postgres"
	"simple-wallet/configs"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// read yaml config file. Dont pass path to read
	// from default path
	cfg := configs.ReadYaml("")
	confg := app.GetConfig(*cfg)

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
