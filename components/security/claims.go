package security

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secretKey")

type Claims struct {
	UserID   string
	IsAdmin  bool
	IsActive bool
	jwt.StandardClaims
}

func (c *Claims) Coder() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(secretKey)
}

// Middleware Function For User
func MiddlewareAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claim, err := ValidateToken(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if !claim.IsAdmin {
			err = errors.New("user not an admin")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if !claim.IsActive {
			err = errors.New("user is not active")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Middleware Function For Contact
func MiddlewareContact(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claim, err := ValidateToken(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if claim.IsAdmin {
			err = errors.New("user Should not be admin")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if !claim.IsActive {
			err = errors.New("user is not active")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Validates Token
func ValidateToken(_ http.ResponseWriter, r *http.Request) (*Claims, error) {

	authCookie, err := r.Cookie("auth")
	tokenString := authCookie.Value
	if err != nil {
		return nil, err
	}

	token, claim, err := checkToken(tokenString)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claim, nil
}

// Checks Token String
func checkToken(tokenString string) (*jwt.Token, *Claims, error) {

	var claim = &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	return token, claim, err
}
