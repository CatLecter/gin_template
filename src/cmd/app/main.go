package main

import (
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"src/app"
	"src/configs"
	"src/database"
	"src/internal/handlers"
	"src/internal/repositories"
	"src/internal/services"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint: true,
	})
	cfg := configs.NewConfig()

	db, err := database.NewPool(
		&cfg.PostgresURI,
		&cfg.MaxConnections,
		&cfg.MinConnections,
		&cfg.MaxConnLifetime,
		&cfg.MaxConnIdleTime,
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err.Error())
	}

	repos := repositories.NewRepository(db)
	service := services.NewService(repos)
	handler := handlers.NewHandler(service)

	server := new(app.App)

	if err := server.Run(cfg.Host, cfg.Port, handler.InitRoutes()); err != nil {
		log.Fatalf("Failed to start server: %v", err.Error())
	}
}
