package main

import (
	"github.com/cheeseNA/owlback/internal/config"
	api "github.com/cheeseNA/owlback/internal/ogen"
	"github.com/cheeseNA/owlback/internal/repository"
	"github.com/cheeseNA/owlback/internal/service"
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
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
