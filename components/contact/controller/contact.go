package controller

import (
	"contact-app-main/components/contact/service"
	"contact-app-main/components/security"
	"contact-app-main/components/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UserInput struct {
	FName    string `json:"first_name"`
	LName    string `json:"last_name"`
	IsActive bool   `json:"is_active"`
}

func CreateContact(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userIdFromURL := vars["userId"]

	var userInput UserInput

	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	claims, err := security.ValidateToken(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if userIdFromURL != claims.UserID {
		http.Error(w, "userId mismatch", http.StatusForbidden)
		return
	}

	newContact, err := service.CreateContact(userIdFromURL, userInput.FName, userInput.LName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, newContact)
}

func GetAllContacts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdFromURL := vars["userId"]

	claims, err := security.ValidateToken(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	if claims.UserID != userIdFromURL {
		http.Error(w, "Forbidden: userId mismatch", http.StatusForbidden)
		return
	}

	contacts, err := service.GetAllContacts(userIdFromURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, contacts)
}

func GetContactByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdFromURL := vars["userId"]
	contactId := vars["contactId"]

	claims, err := security.ValidateToken(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if claims.UserID != userIdFromURL {
		http.Error(w, "userId mismatch", http.StatusForbidden)
		return
	}

	contact, err := service.GetContactByID(userIdFromURL, contactId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.WriteJSON(w, http.StatusOK, contact)
}

func UpdateContactById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdFromURL := vars["userId"]
	contactId := vars["contactId"]

	var input UserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid json body", http.StatusBadRequest)
		return
	}

	claims, err := security.ValidateToken(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	if claims.UserID != userIdFromURL {
		http.Error(w, "User ID mismatch", http.StatusForbidden)
		return
	}

	updatedContact, err := service.UpdateContactById(userIdFromURL, contactId, input.FName, input.LName, input.IsActive)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	utils.WriteJSON(w, http.StatusOK, updatedContact)
}

func DeleteContactByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdFromURL := vars["userId"]
	contactId := vars["contactId"]

	claims, err := security.ValidateToken(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if claims.UserID != userIdFromURL {
		http.Error(w, "User ID mismatch", http.StatusForbidden)
		return
	}

	err = service.DeleteContactById(userIdFromURL, contactId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
