package controller

import (
	"contact-app-main/components/apperror"
	"contact-app-main/components/contactDetail/service"
	"contact-app-main/components/security"
	"contact-app-main/components/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ContactDetailInput struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

func CreateContactDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	contactId := vars["contactId"]

	var input ContactDetailInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusBadRequest, "InvalidJSON", "Invalid JSON body", err))
		return
	}

	claims, err := security.ValidateToken(w, r)
	if err != nil || claims.UserID != userId {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusUnauthorized, "Unauthorized", "invalid token or userId mismatch", err))
		return
	}

	contactDetail, appErr := service.CreateContactDetail(userId, contactId, input.Type, input.Value)
	if appErr != nil {
		utils.RespondWithAppError(w, appErr)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, contactDetail)
}

func GetAllContactDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	contactId := vars["contactId"]

	claims, err := security.ValidateToken(w, r)
	if err != nil || claims.UserID != userId {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusUnauthorized, "Unauthorized", "invalid token or userId mismatch", err))
		return
	}

	details, appErr := service.GetContactDetails(userId, contactId)
	if appErr != nil {
		utils.RespondWithAppError(w, appErr)
		return
	}

	utils.WriteJSON(w, http.StatusOK, details)
}

func GetContactDetailById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	contactId := vars["contactId"]
	contactDetailId := vars["ContactDetailId"]

	claims, err := security.ValidateToken(w, r)
	if err != nil || claims.UserID != userId {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusUnauthorized, "Unauthorized", "invalid token or userId mismatch", err))
		return
	}

	detail, appErr := service.GetContactDetailById(userId, contactId, contactDetailId)
	if appErr != nil {
		utils.RespondWithAppError(w, appErr)
		return
	}

	utils.WriteJSON(w, http.StatusOK, detail)
}

func UpdateContactDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	contactId := vars["contactId"]
	contactDetailId := vars["contactDetailId"]

	var input ContactDetailInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusBadRequest, "InvalidJSON", "Invalid JSON body", err))
		return
	}

	claims, err := security.ValidateToken(w, r)
	if err != nil || claims.UserID != userId {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusUnauthorized, "Unauthorized", "invalid token or userId mismatch", err))
		return
	}

	updatedDetail, appErr := service.UpdateContactDetail(userId, contactId, contactDetailId, input.Type, input.Value)
	if appErr != nil {
		utils.RespondWithAppError(w, appErr)
		return
	}

	utils.WriteJSON(w, http.StatusOK, updatedDetail)
}

func DeleteContactDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	contactId := vars["contactId"]
	contactDetailId := vars["contactDetailId"]

	claims, err := security.ValidateToken(w, r)
	if err != nil || claims.UserID != userId {
		utils.RespondWithAppError(w, apperror.NewAppError(http.StatusUnauthorized, "Unauthorized", "invalid token or userId mismatch", err))
		return
	}

	appErr := service.DeleteContactDetail(userId, contactId, contactDetailId)
	if appErr != nil {
		utils.RespondWithAppError(w, appErr)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
