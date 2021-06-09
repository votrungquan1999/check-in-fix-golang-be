package models

type Reviews struct {
	ID           *string  `firestore:"id" json:"id"`
	CustomerID   *string  `firestore:"customer_id" json:"customer_id"`
	SubscriberID *string  `firestore:"subscriber_id" json:"subscriber_id"`
	Rating       *float64 `firestore:"rating" json:"rating"`
	Feedback     *string  `firestore:"feedback" json:"feedback"`
	IsReviewed   bool     `firestore:"is_reviewed" json:"is_reviewed"`
}
