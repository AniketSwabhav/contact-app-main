package controller

import (
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
	var userInput *UserInput
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdUser, err := service.CreateAdmin(userInput.FirstName, userInput.LastName, userInput.Email, userInput.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, createdUser)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var userInput *UserInput
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdUser, err := service.CreateUser(userInput.FirstName, userInput.LastName, userInput.Email, userInput.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, createdUser)
}

func Login(w http.ResponseWriter, r *http.Request) {

	var userCredentials *credential.Credentials
	err := json.NewDecoder(r.Body).Decode(&userCredentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validates Credentials
	err = userCredentials.ValidateCredential()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Finds User
	foundUser, err := service.Login(userCredentials.Email, userCredentials.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
		http.Error(w, err.Error(), http.StatusBadGateway)
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

	// id := r.URL.Query().Get("id")
	// if id == "" {
	// 	http.Error(w, "missing 'id' query parameter", http.StatusBadRequest)
	// 	return
	// }

	vars := mux.Vars(r)
	userIdFromURL := vars["id"]

	foundUser, err := service.FindUserByID(userIdFromURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.WriteJSON(w, http.StatusOK, foundUser)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := service.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, users)
}

// Put User
func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	// id := r.URL.Query().Get("id")
	// if id == "" {
	// 	http.Error(w, "missing 'id' query parameter", http.StatusBadRequest)
	// 	return
	// }

	vars := mux.Vars(r)
	userIdFromURL := vars["id"]

	var user *UserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedUser, err := service.Update(userIdFromURL, user.FirstName, user.LastName, user.IsAdmin, user.IsActive)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.WriteJSON(w, http.StatusOK, updatedUser)
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {

	// id := r.URL.Query().Get("id")
	// if id == "" {
	// 	http.Error(w, "missing 'id' query parameter", http.StatusBadRequest)
	// 	return
	// }

	vars := mux.Vars(r)
	userIdFromURL := vars["id"]

	err := service.DeleteUserByID(userIdFromURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	utils.WriteJSON(w, http.StatusNoContent, nil)
}
