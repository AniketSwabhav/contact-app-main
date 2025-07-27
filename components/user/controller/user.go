package controller

import (
	"contact-app-main/components/apperror"
	"contact-app-main/components/security"
	"contact-app-main/components/user/service"
	"contact-app-main/components/utils"
	"contact-app-main/models/credential"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type UserInput struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	IsAdmin   bool
	IsActive  bool
}

func RegisterAdmin(w http.ResponseWriter, r *http.Request) {
	var userInput UserInput
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Failed to parse request body", err))
		return
	}

	if userInput.FirstName == "" || userInput.LastName == "" || userInput.Email == "" || userInput.Password == "" {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusBadRequest, "MISSING_FIELDS", "All fields are required", nil))
		return
	}

	createdUser, appErr := service.CreateAdmin(userInput.FirstName, userInput.LastName, userInput.Email, userInput.Password)
	if appErr != nil {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusInternalServerError, "CREATE_ADMIN_FAILED", "Failed to create admin user", err))
		return
	}

	utils.WriteJSON(w, http.StatusCreated, createdUser)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var userInput UserInput
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Failed to parse request body", err))
		return
	}

	if userInput.FirstName == "" || userInput.LastName == "" || userInput.Email == "" || userInput.Password == "" {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusBadRequest, "MISSING_FIELDS", "All fields are required", nil))
		return
	}

	createdUser, appErr := service.CreateUser(userInput.FirstName, userInput.LastName, userInput.Email, userInput.Password)
	if appErr != nil {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusInternalServerError, "CREATE_USER_FAILED", "Failed to create user", err))
		return
	}

	utils.WriteJSON(w, http.StatusCreated, createdUser)
}

func Login(w http.ResponseWriter, r *http.Request) {

	var userCredentials *credential.Credentials
	err := json.NewDecoder(r.Body).Decode(&userCredentials)
	if err != nil {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Failed to parse request body", err))
		return
	}

	err = userCredentials.ValidateCredential()
	if err != nil {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusBadRequest, "INVALID_CREDENTIALS", err.Error(), nil))
		return
	}

	foundUser, appErr := service.Login(userCredentials.Email, userCredentials.Password)
	if appErr != nil {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusUnauthorized, "LOGIN_FAILED", "Invalid email or password", err))
		return
	}

	// JWT Creation
	claim := &security.Claims{
		UserID:   foundUser.UserID,
		IsAdmin:  foundUser.IsAdmin,
		IsActive: foundUser.IsActive,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
		},
	}

	token, err := claim.Coder()
	if err != nil {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusInternalServerError, "TOKEN_GENERATION_FAILED", "Failed to generate token", err))
		return
	}

	// Sets Cookie named "auth"
	http.SetCookie(w, &http.Cookie{
		Name:  "auth",
		Value: token,
	})

	utils.WriteJSON(w, http.StatusAccepted, foundUser)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userIdFromURL := vars["id"]

	foundUser, err := service.FindUserByID(userIdFromURL)
	if err != nil {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusNotFound, "USER_NOT_FOUND", err.Error(), err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, foundUser)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := service.GetAllUsers()
	if err != nil {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusInternalServerError, "GET_USERS_FAILED", err.Error(), err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, users)
}

// Put User
func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdFromURL := vars["id"]

	var userInput UserInput
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusBadRequest, "INVALID_JSON", "Failed to parse request body", err))
		return
	}

	updatedUser, err := service.Update(userIdFromURL, userInput.FirstName, userInput.LastName, userInput.IsAdmin, userInput.IsActive)
	if err != nil {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusBadRequest, "UPDATE_FAILED", err.Error(), err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, updatedUser)
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userIdFromURL := vars["id"]

	err := service.DeleteUserByID(userIdFromURL)
	if err != nil {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusNotFound, "DELETE_FAILED", err.Error(), err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
