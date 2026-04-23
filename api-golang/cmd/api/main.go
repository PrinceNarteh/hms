package main

import (
	"log"

	"hms-backend/internals/config"
	"hms-backend/internals/db"
	"hms-backend/internals/repository"
)

func main() {
	// connect to database
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	repository.InitRespository(db)

	// initialize fiber app
	app := initApp()

	log.Fatal(app.Listen(config.Envs.AppConfig.Port))
}
