package models

type Tickets struct {
	ID          *string `json:"id,omitempty"`
	CustomerID  *string `firestore:"customer_id,omitempty" json:"customer_id"`
	ServiceID   *string `firestore:"service_id,omitempty" json:"service_id"`
	Description *string `firestore:"description,omitempty" json:"description"`
	ApprovedBy  *string `firestore:"approved_by,omitempty" json:"approved_by"`
	PhoneType   *string `firestore:"phone_type,omitempty" json:"phone_type"`
}

type DraftTickets struct {
	ID          string `firestore:"id,omitempty" json:"id,omitempty"`
	UserID      string `firestore:"user_id,omitempty" json:"user_id,omitempty"`
	ServiceID   int    `firestore:"service_id,omitempty" json:"service_id,omitempty"`
	Description string `firestore:"description,omitempty" json:"description,omitempty"`
}
