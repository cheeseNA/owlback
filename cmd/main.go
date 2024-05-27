package main

import (
	"github.com/cheeseNA/owlback/internal/config"
	api "github.com/cheeseNA/owlback/internal/ogen"
	"github.com/cheeseNA/owlback/internal/repository"
	"github.com/cheeseNA/owlback/internal/service"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	logger := zap.NewExample()
	defer func() {
		err := logger.Sync()
		if err != nil {
			panic(err)
		}
	}()

	logger.Info("Starting server")
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: cfg.PostgresConnectionString,
	}))
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&repository.Task{}); err != nil {
		panic(err)
	}

	repo := repository.NewTaskRepository(db, logger)
	s := service.NewService(repo)
	srv, err := api.NewServer(s)
	if err != nil {
		log.Fatal(err)
	}

	AllowedOrigins := []string{}
	if cfg.RunningEnvironment == config.Local {
		AllowedOrigins = append(AllowedOrigins, "http://localhost:3000")
	} else if cfg.RunningEnvironment == config.Production {
		AllowedOrigins = append(AllowedOrigins, "https://crawl-owl.vercel.app")
	}
	c := cors.New(cors.Options{
		AllowedOrigins:   AllowedOrigins,
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: cfg.RunningEnvironment == config.Local,
	})
	// Insert the middleware
	handler := c.Handler(srv)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
