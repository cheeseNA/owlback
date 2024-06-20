package main

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/cheeseNA/owlback/internal/config"
	"github.com/cheeseNA/owlback/internal/funccall"
	"github.com/cheeseNA/owlback/internal/middleware"
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
	if err := db.AutoMigrate(&repository.Task{}, &repository.User{}); err != nil {
		panic(err)
	}

	repo := repository.NewTaskRepository(db, logger)
	funcService, err := funccall.NewFuncService()
	if err != nil {
		panic(err)
	}
	s := service.NewService(repo, funcService, logger)
	srv, err := api.NewServer(s)
	if err != nil {
		log.Fatal(err)
	}

	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	handler := middleware.Auth(authClient)(srv)

	AllowedOrigins := []string{}
	if cfg.RunningEnvironment == config.Local {
		AllowedOrigins = append(AllowedOrigins, "http://localhost:3000")
	} else if cfg.RunningEnvironment == config.Production {
		AllowedOrigins = append(AllowedOrigins, "https://crawl-owl.vercel.app")
	}
	c := cors.New(cors.Options{
		AllowedOrigins:   AllowedOrigins,
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	})
	// Insert the middleware
	handler = c.Handler(handler)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
