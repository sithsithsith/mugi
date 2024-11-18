package main

import (
	"log"
	"net/http"

	"github.com/sithsithsith/cognito-auth-service/internal/app"
	"github.com/sithsithsith/cognito-auth-service/pkg/logger"
)

func main() {
	logger.Init()
	app := app.NewApp()
	app.RegisterRoutes()

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
