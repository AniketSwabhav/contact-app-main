package service

import (
	"contact-app-main/models/credential"
	"contact-app-main/models/user"
	"errors"
)

func CreateAdmin(fname, lname, email, password string) (*user.User, error) {

	newCredential, err := credential.CreateCredential(email, password)
	if err != nil {
		return nil, err
	}

	user := user.CreateAdmin(fname, lname, email, newCredential.CredentialID)

	return user, nil
}

func CreateUser(fname, lname, email, password string) (*user.User, error) {
	newCredential, err := credential.CreateCredential(email, password)
	if err != nil {
		return nil, err
	}

	user := user.CreateUser(fname, lname, email, newCredential.CredentialID)

	return user, nil
}

func Login(email, password string) (*user.User, error) {
	foundCredentials, err := credential.FindCredential(email)
	if err != nil {
		return nil, err
	}

	if foundCredentials == nil {
		return nil, errors.New("user credentials not found")
	}

	err = credential.CheckPassword(foundCredentials.Password, password)
	if err != nil {
		return nil, err
	}

	foundUser, err := user.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}

func FindUserByID(id string) (*user.User, error) {
	foundUser, err := user.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	if foundUser == nil {
		return nil, errors.New("user not found")
	}

	return foundUser, nil
}

func GetAllUsers() ([]*user.User, error) {
	users := user.GetAllUsers()
	if len(users) == 0 {
		return nil, errors.New("no users found")
	}

	return users, nil
}

func Update(id, fname, lname string, isAdmin, isActive bool) (*user.User, error) {

	existingUser, err := user.FindUserByID(id)
	if err != nil {
		return nil, err
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

func DeleteUserByID(id string) error {
	err := user.DeleteUserByID(id)
	if err != nil {
		return err
	}

	return nil
}
