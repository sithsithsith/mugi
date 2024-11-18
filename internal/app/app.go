package app

import (
	"net/http"

	"github.com/sithsithsith/cognito-auth-service/internal/config"
	"github.com/sithsithsith/cognito-auth-service/internal/handlers"
	"github.com/sithsithsith/cognito-auth-service/internal/services"
)

type App struct {
	CognitoService  *services.CognitoService
	DatabaseService *services.DatabaseService
}

func NewApp() *App {
	client := config.GetCognitoClient()
	cognitoService := services.NewCognitoService(client)
	databaseService := services.NewDatabaseService()

	return &App{
		CognitoService:  cognitoService,
		DatabaseService: databaseService,
	}
}

func (a *App) RegisterRoutes() {
	authHandler := handlers.NewAuthHandler(a.CognitoService, a.DatabaseService)

	http.HandleFunc("/signup", authHandler.SignUpHandler)
	http.HandleFunc("/confirm-signup", authHandler.ConfirmSignUpHandler)
	http.HandleFunc("/signin", authHandler.SignInHandler)
	http.HandleFunc("/migrate", authHandler.MigrateUsersHandler)
}
