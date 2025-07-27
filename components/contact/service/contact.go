package service

import (
	"contact-app-main/components/apperror"
	"contact-app-main/models/contact"
	"contact-app-main/models/user"
	"net/http"
)

func CreateContact(userId, fName, lName string) (*contact.Contact, *apperror.AppError) {

	targetUser, err := user.FindUserByID(userId)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusNotFound, "UserNotFound", "user not found", err)
	}

	createdContact, err := targetUser.CreateContact(fName, lName)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusBadRequest, "InvalidInput", "failed to create contact", err)
	}

	return createdContact, nil

}

func GetAllContacts(userId string) ([]*contact.Contact, *apperror.AppError) {

	targetUser, err := user.FindUserByID(userId)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusNotFound, "UserNotFound", "user not found", err)
	}
	return targetUser.Contacts, nil
}

func GetContactByID(userId, contactId string) (*contact.Contact, *apperror.AppError) {
	targetUser, err := user.FindUserByID(userId)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusNotFound, "UserNotFound", "user not found", err)
	}

	targetContact, err := targetUser.GetContactByID(contactId)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusNotFound, "ContactNotFound", "contact not found", err)
	}

	return targetContact, nil
}

func UpdateContactById(userId, contactId, fName, lName string, isActive bool) (*contact.Contact, error) {
	targetUser, err := user.FindUserByID(userId)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusNotFound, "UserNotFound", "user not found", err)
	}

	updatedContact, err := targetUser.UpdateContactById(contactId, fName, lName, isActive)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusBadRequest, "InvalidInput", "failed to update contact", err)
	}

	return updatedContact, nil

}

func DeleteContactById(userId, contactId string) error {

	targetUser, err := user.FindUserByID(userId)
	if err != nil {
		return apperror.NewAppError(http.StatusNotFound, "UserNotFound", "user not found", err)
	}

	err = targetUser.DeleteContactByID(contactId)
	if err != nil {
		return apperror.NewAppError(http.StatusNotFound, "ContactNotFound", "contact not found", err)
	}

	return nil
}
