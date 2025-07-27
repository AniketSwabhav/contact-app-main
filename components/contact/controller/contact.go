package controller

import (
	"contact-app-main/components/apperror"
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
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusBadRequest, "InvalidJSON", "Invalid JSON body", err))
		return
	}

	claims, err := security.ValidateToken(w, r)
	if err != nil {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusUnauthorized, "Unauthorized", err.Error(), nil))
		return
	}

	if userIdFromURL != claims.UserID {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusForbidden, "Forbidden", "userId mismatch", nil))
		return
	}

	newContact, appErr := service.CreateContact(userIdFromURL, userInput.FName, userInput.LName)
	if appErr != nil {
		utils.RespondWithAppError(w, appErr)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, newContact)
}

func GetAllContacts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdFromURL := vars["userId"]

	claims, err := security.ValidateToken(w, r)
	if err != nil {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusUnauthorized, "Unauthorized", err.Error(), nil))
		return
	}

	if claims.UserID != userIdFromURL {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusForbidden, "Forbidden", "userId mismatch", nil))
		return
	}

	contacts, appErr := service.GetAllContacts(userIdFromURL)
	if appErr != nil {
		utils.RespondWithAppError(w, appErr)
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
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusUnauthorized, "Unauthorized", err.Error(), nil))
		return
	}

	if claims.UserID != userIdFromURL {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusForbidden, "Forbidden", "userId mismatch", nil))
		return
	}

	contact, appErr := service.GetContactByID(userIdFromURL, contactId)
	if appErr != nil {
		utils.RespondWithAppError(w, appErr)
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
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusBadRequest, "InvalidJSON", "Invalid JSON body", err))
		return
	}

	claims, err := security.ValidateToken(w, r)
	if err != nil {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusUnauthorized, "Unauthorized", err.Error(), nil))
	}

	if claims.UserID != userIdFromURL {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusForbidden, "Forbidden", "userId mismatch", nil))
		return
	}

	updatedContact, err := service.UpdateContactById(userIdFromURL, contactId, input.FName, input.LName, input.IsActive)
	if err != nil {
		if appErr, ok := err.(*apperror.AppError); ok {
			utils.RespondWithAppError(w, appErr)
		} else {
			utils.RespondWithAppError(w, apperror.NewAppError(http.StatusInternalServerError, "InternalError", err.Error(), err))
		}
		return
	}

	utils.WriteJSON(w, http.StatusOK, updatedContact)
}

func DeleteContactByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdFromURL := vars["userId"]
	contactId := vars["contactId"]

	claims, err := security.ValidateToken(w, r)
	if err != nil {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusUnauthorized, "Unauthorized", err.Error(), nil))
		return
	}

	if claims.UserID != userIdFromURL {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusForbidden, "Forbidden", "userId mismatch", nil))
		return
	}

	err = service.DeleteContactById(userIdFromURL, contactId)
	if err != nil {
		if appErr, ok := err.(*apperror.AppError); ok {
			utils.RespondWithAppError(w, appErr)
		} else {
			utils.RespondWithAppError(w, apperror.NewAppError(http.StatusInternalServerError, "InternalError", err.Error(), err))
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
