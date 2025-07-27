package middleware

import (
	"contact-app-main/components/apperror"
	"contact-app-main/components/security"
	"contact-app-main/components/utils"
	"net/http"
)

func MiddlewareAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claim, err := security.ValidateToken(w, r)
		if err != nil {
			appErr := apperror.NewAppError(http.StatusUnauthorized, "Unauthorized", "Invalid or missing token", err)
			utils.RespondWithAppError(w, appErr)
			return
		}

		if !claim.IsAdmin {
			appErr := apperror.NewAppError(http.StatusForbidden, "AccessDenied", "User is not an admin", nil)
			utils.RespondWithAppError(w, appErr)
			return
		}

		if !claim.IsActive {
			appErr := apperror.NewAppError(http.StatusUnauthorized, "InactiveUser", "User is not active", nil)
			utils.RespondWithAppError(w, appErr)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func MiddlewareContact(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claim, err := security.ValidateToken(w, r)
		if err != nil {
			appErr := apperror.NewAppError(http.StatusUnauthorized, "Unauthorized", "Invalid or missing token", err)
			utils.RespondWithAppError(w, appErr)
			return
		}
		if claim.IsAdmin {
			appErr := apperror.NewAppError(http.StatusUnauthorized, "AccessDenied", "Current User is an admin", nil)
			utils.RespondWithAppError(w, appErr)
			return
		}
		if !claim.IsActive {
			appErr := apperror.NewAppError(http.StatusUnauthorized, "InactiveUser", "User is not active", nil)
			utils.RespondWithAppError(w, appErr)
			return
		}
		next.ServeHTTP(w, r)
	})
}
