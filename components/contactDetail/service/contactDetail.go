package service

import (
	"contact-app-main/components/apperror"
	contactdetail "contact-app-main/models/contactDetail"
	"contact-app-main/models/user"
	"net/http"
)

func CreateContactDetail(userId, contactId, detailType string, value interface{}) (*contactdetail.ContactDetail, *apperror.AppError) {

	targetUser, err := user.FindUserByID(userId)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusNotFound, "UserNotFound", "User not found", err)
	}

	targetContact, err := targetUser.GetContactByID(contactId)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusNotFound, "ContactNotFound", "Contact not found", err)
	}

	newDetail, err := targetContact.CreateNewContactDetail(detailType, value)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusBadRequest, "InvalidInput", "Failed to create contact detail", err)
	}

	return newDetail, nil
}

func GetContactDetails(userId, contactId string) ([]*contactdetail.ContactDetail, *apperror.AppError) {
	targetUser, err := user.FindUserByID(userId)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusNotFound, "UserNotFound", "User not found", err)
	}

	targetContact, err := targetUser.GetContactByID(contactId)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusNotFound, "ContactNotFound", "Contact not found", err)
	}

	details, err := targetContact.GetAllContactDetails()
	if err != nil {
		return nil, apperror.NewAppError(http.StatusNotFound, "DetailsNotFound", "No contact details found", err)
	}

	return details, nil
}

func GetContactDetailById(userId, contactId, contactDetailId string) (*contactdetail.ContactDetail, *apperror.AppError) {
	targetUser, err := user.FindUserByID(userId)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusNotFound, "UserNotFound", "User not found", err)
	}

	targetContact, err := targetUser.GetContactByID(contactId)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusNotFound, "ContactNotFound", "Contact not found", err)
	}

	detail, err := targetContact.GetContactDetailById(contactDetailId)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusNotFound, "DetailNotFound", "Contact detail not found", err)
	}

	return detail, nil
}

func UpdateContactDetail(userId, contactId, contactDetailId, newType string, newValue interface{}) (*contactdetail.ContactDetail, *apperror.AppError) {
	targetUser, err := user.FindUserByID(userId)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusNotFound, "UserNotFound", "User not found", err)
	}

	targetContact, err := targetUser.GetContactByID(contactId)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusNotFound, "ContactNotFound", "Contact not found", err)
	}

	updatedDetail, err := targetContact.UpdateContactDetailById(contactDetailId, newType, newValue)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusBadRequest, "UpdateFailed", "Failed to update contact detail", err)
	}

	return updatedDetail, nil
}

func DeleteContactDetail(userId, contactId, contactDetailId string) *apperror.AppError {
	targetUser, err := user.FindUserByID(userId)
	if err != nil {
		return apperror.NewAppError(http.StatusNotFound, "UserNotFound", "User not found", err)
	}

	targetContact, err := targetUser.GetContactByID(contactId)
	if err != nil {
		return apperror.NewAppError(http.StatusNotFound, "ContactNotFound", "Contact not found", err)
	}

	err = targetContact.DeleteContactDetailById(contactDetailId)
	if err != nil {
		return apperror.NewAppError(http.StatusNotFound, "DetailNotFound", "Failed to delete contact detail", err)
	}
	return nil
}
