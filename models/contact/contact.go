package contact

import (
	contactdetail "contact-app-main/models/contactDetail"
	"fmt"

	"github.com/google/uuid"
)

type Contact struct {
	ContactID      string                         `json:"contact_id"`
	FName          string                         `json:"first_name"`
	LName          string                         `json:"last_name"`
	IsActive       bool                           `json:"is_active"`
	ContactDetails []*contactdetail.ContactDetail `json:"contact_details"`
}

func NewContact(fName, lName string) *Contact {

	contactId := uuid.New()
	contact := &Contact{
		ContactID:      contactId.String(),
		FName:          fName,
		LName:          lName,
		IsActive:       true,
		ContactDetails: []*contactdetail.ContactDetail{},
	}

	fmt.Printf("Contact created with ContactId: %s\n", contactId)
	return contact
}

func (c *Contact) CreateNewContactDetail(detailType string, value interface{}) (*contactdetail.ContactDetail, error) {

	if detailType == "" || value == "" {
		return nil, fmt.Errorf("detailType and value must not be empty")
	}

	contactDetail := contactdetail.NewContactDetail(detailType, value)

	c.ContactDetails = append(c.ContactDetails, contactDetail)

	return contactDetail, nil
}

func (c *Contact) GetAllContactDetails() ([]*contactdetail.ContactDetail, error) {

	if c.ContactDetails == nil {
		return nil, fmt.Errorf("no contact details found")
	}
	return c.ContactDetails, nil
}

func (c *Contact) GetContactDetailById(contactDetailId string) (*contactdetail.ContactDetail, error) {

	for _, detail := range c.ContactDetails {
		if detail.ContactDetailID == contactDetailId {
			return detail, nil
		}
	}
	return nil, fmt.Errorf("contact detail with ID %s not found", contactDetailId)
}

func (c *Contact) UpdateContactDetailById(contactDetailId string, newType string, newValue interface{}) (*contactdetail.ContactDetail, error) {
	for _, detail := range c.ContactDetails {
		if detail.ContactDetailID == contactDetailId {
			if newType != "" {
				detail.Type = newType
			}
			if newValue != nil {
				detail.Value = newValue
			}
			return detail, nil
		}
	}
	return nil, fmt.Errorf("contact detail with ID %s not found", contactDetailId)
}

func (c *Contact) DeleteContactDetailById(contactDetailId string) error {
	for i, detail := range c.ContactDetails {
		if detail.ContactDetailID == contactDetailId {
			// Remove from slice
			c.ContactDetails = append(c.ContactDetails[:i], c.ContactDetails[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("contact detail with ID %s not found", contactDetailId)
}
