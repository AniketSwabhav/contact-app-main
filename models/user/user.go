package user

import (
	"contact-app-main/models/contact"
	"fmt"

	"github.com/google/uuid"
)

var userMap = make(map[string]*User)

type User struct {
	UserID       string             `json:"user_id"`
	FName        string             `json:"first_name"`
	LName        string             `json:"last_name"`
	Email        string             `json:"email"`
	IsAdmin      bool               `json:"is_admin"`
	IsActive     bool               `json:"is_active"`
	Contacts     []*contact.Contact `json:"-"`
	CredentialID string             `json:"-"`
}

// New Factory for User
func newUser(fName, lName, email string, isAdmin bool, credentialId string) *User {
	if fName == "" || lName == "" {
		return nil
	}

	id := uuid.New()
	user := &User{
		UserID:       id.String(),
		FName:        fName,
		LName:        lName,
		Email:        email,
		IsAdmin:      isAdmin,
		IsActive:     true,
		Contacts:     []*contact.Contact{},
		CredentialID: credentialId,
	}

	userMap[user.UserID] = user

	role := "Staff"
	if user.IsAdmin {
		role = "Admin"
	}
	fmt.Printf("User Created with UserId : %s (%s)\n", user.UserID, role)
	return user
}

func CreateAdmin(fName, lName, email string, credentialId string) *User {
	newAdmin := newUser(fName, lName, email, true, credentialId)
	return newAdmin
}

func CreateUser(fName, lName, email string, credentialId string) *User {
	newStaff := newUser(fName, lName, email, false, credentialId)
	return newStaff
}

func FindUserByID(id string) (*User, error) {
	user, exists := userMap[id]
	if !exists {
		return nil, fmt.Errorf("user with id %s not found", id)
	}
	return user, nil
}

func GetAllUsers() []*User {
	if len(userMap) == 0 {
		return nil
	}

	users := make([]*User, 0, len(userMap))
	for _, user := range userMap {
		users = append(users, user)
	}
	return users
}

// Finds User by Email
func FindUserByEmail(email string) (*User, error) {
	for _, user := range userMap {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user with email %s not found", email)
}

func DeleteUserByID(id string) error {
	user, exists := userMap[id]
	if !exists {
		return fmt.Errorf("user with id %s not found", id)
	}

	delete(userMap, id)
	fmt.Printf("User with UserId : %s deleted successfully\n", user.UserID)
	return nil
}

func (u *User) CreateContact(fName, lName string) (*contact.Contact, error) {

	if fName == "" || lName == "" {
		return nil, fmt.Errorf("first name and last name are required")
	}

	newContact := contact.NewContact(fName, lName)
	if newContact == nil {
		return nil, fmt.Errorf("failed to create contact")
	}
	u.Contacts = append(u.Contacts, newContact)
	return newContact, nil
}

func (u *User) GetContactByID(contactId string) (*contact.Contact, error) {
	for _, c := range u.Contacts {
		if c.ContactID == contactId {
			return c, nil
		}
	}
	return nil, fmt.Errorf("contact with ID %s not found", contactId)
}

func (u *User) UpdateContactById(contactId, fName, lName string, isActive bool) (*contact.Contact, error) {
	for _, c := range u.Contacts {
		if c.ContactID == contactId {
			c.FName = fName
			c.LName = lName
			c.IsActive = isActive
			return c, nil
		}
	}
	return nil, fmt.Errorf("contact with ID %s not found", contactId)
}

func (u *User) DeleteContactByID(contactId string) error {
	for i, c := range u.Contacts {
		if c.ContactID == contactId {
			u.Contacts = append(u.Contacts[:i], u.Contacts[i+1:]...)
			return nil
		}
	}
	return nil
}
