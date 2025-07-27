package credential

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const cost = 10

// to store a map of credentials
var credentialStore = make(map[string]*Credentials)

type Credentials struct {
	CredentialID string
	Email        string
	Password     string
}

func (c *Credentials) ValidateCredential() error {

	if len(strings.TrimSpace(c.Email)) == 0 || len(strings.TrimSpace(c.Password)) == 0 {
		return errors.New("invalid credentials")
	}

	return nil
}

func CreateCredential(email string, password string) (*Credentials, error) {

	if len(strings.TrimSpace(email)) == 0 || len(strings.TrimSpace(password)) == 0 {
		return nil, errors.New("credentials cannot be empty")
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	credentialId := uuid.New()

	newCredential := &Credentials{
		CredentialID: credentialId.String(),
		Email:        email,
		Password:     string(hashedPassword),
	}

	credentialStore[credentialId.String()] = newCredential

	return newCredential, nil
}

// Hashes Password
func hashPassword(password string) ([]byte, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return nil, err
	}

	return hashedPassword, nil
}

// Finds Credential by Email
func FindCredential(email string) (*Credentials, error) {

	for _, credential := range credentialStore {
		if credential.Email == email {
			return credential, nil
		}
	}
	return nil, errors.New("credential not found")
}

// Checks Password
func CheckPassword(userPassword string, inputPassword string) error {

	return bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(inputPassword))

}

func UpdatePassword(credentialId string, newPassword string) error {

	credential, exists := credentialStore[credentialId]
	if !exists {
		return errors.New("credential not found")
	}

	hashedPassword, err := hashPassword(newPassword)
	if err != nil {
		return err
	}

	credential.Password = string(hashedPassword)
	return nil
}
