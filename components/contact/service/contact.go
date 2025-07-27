package service

import (
	"contact-app-main/models/contact"
	"contact-app-main/models/user"
)

func CreateContact(userId, fName, lName string) (*contact.Contact, error) {

	targetUser, err := user.FindUserByID(userId)
	if err != nil {
		return nil, err
	}

	createdContact, err := targetUser.CreateContact(fName, lName)
	if err != nil {
		return nil, err
	}

	return createdContact, nil

}

func GetAllContacts(userId string) ([]*contact.Contact, error) {
	targetUser, err := user.FindUserByID(userId)
	if err != nil {
		return nil, err
	}

	return targetUser.Contacts, nil
}

func GetContactByID(userId, contactId string) (*contact.Contact, error) {
	targetUser, err := user.FindUserByID(userId)
	if err != nil {
		return nil, err
	}

	targetContact, err := targetUser.GetContactByID(contactId)

	return targetContact, err
}

func UpdateContactById(userId, contactId, fName, lName string, isActive bool) (*contact.Contact, error) {

	targetUser, err := user.FindUserByID(userId)
	if err != nil {
		return nil, err
	}

	updatedContact, err := targetUser.UpdateContactById(contactId, fName, lName, isActive)
	if err != nil {
		return nil, err
	}

	return updatedContact, nil

}

func DeleteContactById(userId, contactId string) error {

	targetUser, err := user.FindUserByID(userId)
	if err != nil {
		return err
	}

	targetUser.DeleteContactByID(contactId)

	return nil
}
