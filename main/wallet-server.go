package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber"

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

	//userRepo := user.NewRepository(database)
	//u, err := userRepo.GetByPhoneNumber("world@helloo.com")
	//if err != nil {
	//	log.Printf("%s", err)
	//	os.Exit(1)
	//}
	//log.Println(u)

	domain := registry.NewDomain(confg, database)

	// create the fiber server.
	server := fiber.New()
	rest.Router(server, domain) // add endpoints

	// listen and serve
	port := fmt.Sprintf(":%v", 6700)
	log.Fatal(server.Listen(port))
}
