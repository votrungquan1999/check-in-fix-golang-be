package models

type Customers struct {
	ID           *string `json:"id,omitempty"`
	PhoneNumber  *string `firestore:"phone_number,omitempty" json:"phone_number"`
	FirstName    *string `firestore:"first_name,omitempty" json:"first_name"`
	LastName     *string `firestore:"last_name,omitempty" json:"last_name"`
	Email        *string `firestore:"email,omitempty" json:"email"`
	SubscriberID *string `firestore:"subscriber_id,omitempty" json:"subscriber_id"`
}
