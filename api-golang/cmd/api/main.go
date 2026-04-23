package main

import (
	"log"

	"hms-backend/internals/config"
	"hms-backend/internals/db"
	"hms-backend/internals/repository"
	"hms-backend/internals/services"
)

func main() {
	// connect to database
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// initialize repository
	repo := repository.InitRespository(db)

	// initialize services
	services.InitServices(repo)

	// initialize fiber app
	app := initApp()

	log.Fatal(app.Listen(config.Envs.AppConfig.Port))
}
