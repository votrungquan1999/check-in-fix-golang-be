package models

type RatingPlatform struct {
	ID           *string `firestore:"id" json:"id"`
	SubscriberID *string `firestore:"subscriber_id" json:"subscriber_id"`
	Name         *string `firestore:"name" json:"name"`
	URL          *string `firestore:"url" json:"url"`
}
