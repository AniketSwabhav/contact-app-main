package service

import (
	"contact-app-main/components/apperror"
	"contact-app-main/models/credential"
	"contact-app-main/models/user"
	"net/http"
)

func CreateAdmin(fname, lname, email, password string) (*user.User, *apperror.AppError) {

	newCredential, err := credential.CreateCredential(email, password)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusInternalServerError, "CREDENTIAL_CREATION_FAILED", "Failed to create credential", err)
	}

	user := user.CreateAdmin(fname, lname, email, newCredential.CredentialID)
	if user == nil {
		return nil, apperror.NewAppError(http.StatusBadRequest, "USER_CREATION_FAILED", "Invalid admin data", nil)
	}

	return user, nil
}

func CreateUser(fname, lname, email, password string) (*user.User, *apperror.AppError) {
	newCredential, err := credential.CreateCredential(email, password)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusInternalServerError, "CREDENTIAL_CREATION_FAILED", "Failed to create credential", err)
	}

	user := user.CreateUser(fname, lname, email, newCredential.CredentialID)
	if user == nil {
		return nil, apperror.NewAppError(http.StatusBadRequest, "USER_CREATION_FAILED", "Invalid user data", nil)
	}

	return user, nil
}

func Login(email, password string) (*user.User, *apperror.AppError) {
	foundCredentials, err := credential.FindCredential(email)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusNotFound, "CREDENTIAL_NOT_FOUND", "Credential not found", err)
	}

	if foundCredentials == nil {
		return nil, apperror.NewAppError(http.StatusUnauthorized, "INVALID_CREDENTIALS", "Invalid email or password", nil)
	}

	if err := credential.CheckPassword(foundCredentials.Password, password); err != nil {
		return nil, apperror.NewAppError(http.StatusUnauthorized, "INVALID_PASSWORD", "Incorrect password", err)
	}

	foundUser, err := user.FindUserByEmail(email)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusNotFound, "USER_NOT_FOUND", "User not found", err)
	}

	return foundUser, nil
}

func FindUserByID(id string) (*user.User, *apperror.AppError) {

	foundUser, err := user.FindUserByID(id)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusNotFound, "USER_NOT_FOUND", "User not found", err)
	}

	return foundUser, nil
}

func GetAllUsers() ([]*user.User, *apperror.AppError) {
	users := user.GetAllUsers()
	if len(users) == 0 {
		return nil, apperror.NewAppError(http.StatusNotFound, "NO_USERS_FOUND", "No users available", nil)
	}
	return users, nil
}

func Update(id, fname, lname string, isAdmin, isActive bool) (*user.User, error) {
	existingUser, err := user.FindUserByID(id)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusNotFound, "USER_NOT_FOUND", "User not found", err)
	}

	if fname != "" {
		existingUser.FName = fname
	}
	if lname != "" {
		existingUser.LName = lname
	}

	existingUser.IsAdmin = isAdmin
	existingUser.IsActive = isActive

	return existingUser, nil
}

func DeleteUserByID(id string) *apperror.AppError {
	err := user.DeleteUserByID(id)
	if err != nil {
		return apperror.NewAppError(http.StatusNotFound, "USER_DELETE_FAILED", "Could not delete user", err)
	}
	return nil
}
