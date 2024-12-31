package main

import (
	"log/slog"
	"net/http"
	"os"
	"splitwise/config"
	"splitwise/handler"
	"splitwise/repository"
	"splitwise/service"
)

const (
	VERSION string = "v1"
	PORT    string = ":8080"
)

func main() {
	db, err := config.InitDb()
	if err != nil {
		slog.Error("error connecting to database", "error", err)
		os.Exit(1)
	}

	slog.Info("connection to db successful")
	slog.Info("server started successfully", "port", PORT)

	userRepository := repository.NewUserRepositoryImpl(db)
	authService := service.NewAuthService(userRepository)
	authHandler := handler.NewAuthHandler(authService)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/{VERSION}/auth/signup", authHandler.Signup)
	mux.HandleFunc("/api/{VERSION}/auth/login", authHandler.Login)
	mux.HandleFunc("/api/{VERSION}/auth/logout", authHandler.Logout)

	if err := http.ListenAndServe(PORT, mux); err != nil {
		slog.Error("error starting server", "error", err)
		os.Exit(1)
	}

}