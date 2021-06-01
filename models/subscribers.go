package models

type Subscribers struct {
	ID    *string `firestore:"id,omitempty" json:"id,omitempty"`
	Email *string `firestore:"email,omitempty" json:"email,omitempty"`
	Name  *string `firestore:"name,omitempty" json:"name,omitempty"`
}
