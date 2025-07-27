package controller

import (
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
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	claims, err := security.ValidateToken(w, r)
	if err != nil || claims.UserID != userId {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	contactDetail, err := service.CreateContactDetail(userId, contactId, input.Type, input.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	details, err := service.GetContactDetails(userId, contactId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
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
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	detail, err := service.GetContactDetailById(userId, contactId, contactDetailId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
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
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	claims, err := security.ValidateToken(w, r)
	if err != nil || claims.UserID != userId {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	updatedDetail, err := service.UpdateContactDetail(userId, contactId, contactDetailId, input.Type, input.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
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
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = service.DeleteContactDetail(userId, contactId, contactDetailId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
