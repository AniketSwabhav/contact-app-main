package service

import (
	contactdetail "contact-app-main/models/contactDetail"
	"contact-app-main/models/user"
)

func CreateContactDetail(userId, contactId, detailType string, value interface{}) (*contactdetail.ContactDetail, error) {

	targetUser, err := user.FindUserByID(userId)
	if err != nil {
		return nil, err
	}

	targetContact, err := targetUser.GetContactByID(contactId)
	if err != nil {
		return nil, err
	}

	newContactDetail, err := targetContact.CreateNewContactDetail(detailType, value)
	if err != nil {
		return nil, err
	}

	return newContactDetail, nil
}

func GetContactDetails(userId, contactId string) ([]*contactdetail.ContactDetail, error) {
	targetUser, err := user.FindUserByID(userId)
	if err != nil {
		return nil, err
	}

	targetContact, err := targetUser.GetContactByID(contactId)
	if err != nil {
		return nil, err
	}

	allDetails, err := targetContact.GetAllContactDetails()
	if err != nil {
		return nil, err
	}

	return allDetails, nil
}

func GetContactDetailById(userId, contactId, contactDetailId string) (*contactdetail.ContactDetail, error) {
	targetUser, err := user.FindUserByID(userId)
	if err != nil {
		return nil, err
	}

	targetContact, err := targetUser.GetContactByID(contactId)
	if err != nil {
		return nil, err
	}

	detail, err := targetContact.GetContactDetailById(contactDetailId)
	if err != nil {
		return nil, err
	}

	return detail, nil
}

func UpdateContactDetail(userId, contactId, contactDetailId, newType string, newValue interface{}) (*contactdetail.ContactDetail, error) {
	targetUser, err := user.FindUserByID(userId)
	if err != nil {
		return nil, err
	}

	targetContact, err := targetUser.GetContactByID(contactId)
	if err != nil {
		return nil, err
	}

	updatedDetail, err := targetContact.UpdateContactDetailById(contactDetailId, newType, newValue)
	if err != nil {
		return nil, err
	}

	return updatedDetail, nil
}

func DeleteContactDetail(userId, contactId, contactDetailId string) error {
	targetUser, err := user.FindUserByID(userId)
	if err != nil {
		return err
	}

	targetContact, err := targetUser.GetContactByID(contactId)
	if err != nil {
		return err
	}

	return targetContact.DeleteContactDetailById(contactDetailId)
}
