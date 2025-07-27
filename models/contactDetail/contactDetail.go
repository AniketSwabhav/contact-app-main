package contactdetail

import (
	"github.com/google/uuid"
)

type ContactDetail struct {
	ContactDetailID string      `json:"contact_detail_id"`
	Type            string      `json:"type"`
	Value           interface{} `json:"value"`
}

func NewContactDetail(detailType string, value interface{}) *ContactDetail {

	contactDetailId := uuid.New()
	contactDetail := &ContactDetail{
		ContactDetailID: contactDetailId.String(),
		Type:            detailType,
		Value:           value,
	}
	return contactDetail
}
