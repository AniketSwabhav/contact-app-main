package middleware

import (
	"contact-app-main/components/apperror"
	"contact-app-main/components/utils"
	"net/http"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if rec := recover(); rec != nil {
				var appErr *apperror.AppError

				switch e := rec.(type) {
				case *apperror.AppError:
					appErr = e
				case error:
					appErr = apperror.NewAppError(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "An unexpected error occurred", e)
				default:
					appErr = apperror.NewAppError(http.StatusInternalServerError, "UNKNOWN_ERROR", "An unknown error occurred", nil)
				}

				utils.WriteJSON(w, appErr.StatusCode, appErr)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
