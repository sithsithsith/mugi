package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sithsithsith/cognito-auth-service/internal/services"
	"github.com/sithsithsith/cognito-auth-service/internal/utils"
)

type AuthHandler struct {
	CognitoService  *services.CognitoService
	DatabaseService *services.DatabaseService
}

func NewAuthHandler(cs *services.CognitoService, ds *services.DatabaseService) *AuthHandler {
	return &AuthHandler{
		CognitoService:  cs,
		DatabaseService: ds,
	}
}

func (ah *AuthHandler) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var req map[string]string
	json.NewDecoder(r.Body).Decode(&req)

	err := ah.CognitoService.SignUp(req["phone_number"], req["password"])
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Sign-up successful")
}

func (ah *AuthHandler) ConfirmSignUpHandler(w http.ResponseWriter, r *http.Request) {
	var req map[string]string
	json.NewDecoder(r.Body).Decode(&req)

	err := ah.CognitoService.ConfirmSignUp(req["phone_number"], req["code"])
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Confirmation successful")
}

func (ah *AuthHandler) SignInHandler(w http.ResponseWriter, r *http.Request) {
	var req map[string]string
	json.NewDecoder(r.Body).Decode(&req)

	token, err := ah.CognitoService.SignIn(req["phone_number"], req["password"])
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (ah *AuthHandler) MigrateUsersHandler(w http.ResponseWriter, r *http.Request) {
	migrationService := services.NewMigrationService(ah.CognitoService, ah.DatabaseService)

	err := migrationService.MigrateUsers()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Migration failed")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Users migrated successfully")
}
