package utils

import (
	"contact-app-main/components/apperror"
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func RespondWithAppError(w http.ResponseWriter, appErr *apperror.AppError) {
	WriteJSON(w, appErr.StatusCode, appErr)
}
