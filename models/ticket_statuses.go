package models

type TicketStatuses struct {
	ID           *string `firestore:"id" json:"id"`
	Name         *string `firestore:"name" json:"name"`
	SubscriberID *string `firestore:"subscriber_id" json:"subscriber_id"`
	Order        *int64  `firestore:"order" json:"order"`
	Type         *int64  `firestore:"type" json:"type"`
}
