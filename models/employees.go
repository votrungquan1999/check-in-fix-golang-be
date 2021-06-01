package models

type Employees struct {
	ID           *string  `firestore:"id,omitempty" json:"id,omitempty"`
	UserID       *string  `firestore:"user_id,omitempty" json:"user_id" binding:"omitempty"`
	Email        *string  `firestore:"email,omitempty" json:"email" binding:"omitempty"`
	FirstName    *string  `firestore:"first_name,omitempty" json:"first_name" binding:"omitempty"`
	LastName     *string  `firestore:"last_name,omitempty" json:"last_name" binding:"omitempty"`
	SubscriberID *string  `firestore:"subscriber_id,omitempty" json:"subscriber_id" binding:"omitempty"`
	Scopes       []string `firestore:"scope,omitempty" json:"scopes" binding:"omitempty"`
}
