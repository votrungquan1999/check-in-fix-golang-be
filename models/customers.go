package models

type Customers struct {
	ID           *string `json:"id,omitempty"`
	PhoneNumber  *string `firestore:"phone_number,omitempty" json:"phone_number"`
	FirstName    *string `firestore:"first_name,omitempty" json:"first_name"`
	LastName     *string `firestore:"last_name,omitempty" json:"last_name"`
	Email        *string `firestore:"email,omitempty" json:"email"`
	SubscriberID *string `firestore:"subscriber_id,omitempty" json:"subscriber_id"`
	AddressLine1 *string `firestore:"address_line_1,omitempty" json:"address_line_1"`
	AddressLine2 *string `firestore:"address_line_2,omitempty" json:"address_line_2"`
	City         *string `firestore:"city,omitempty" json:"city"`
	State        *string `firestore:"state,omitempty" json:"state"`
	ZipCode      *string `firestore:"zip_code,omitempty" json:"zip_code"`
	Country      *string `firestore:"country,omitempty" json:"country"`
}
