package main

import (
	"log"

	"hms-backend/internals/config"
	"hms-backend/internals/db"
	"hms-backend/internals/handlers"
	"hms-backend/internals/repository"
	"hms-backend/internals/routes"
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
	svc := services.InitServices(repo)

	// initialize handlers
	handlers := handlers.InitHandlers(svc)

	// initialize fiber app
	app := initApp()

	routes.InitRoutes(app, handlers)

	log.Fatal(app.Listen(config.Envs.AppConfig.Port))
}
