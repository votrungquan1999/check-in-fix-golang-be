package requests

type SMSSendingRequest struct {
	CustomerIds []string `json:"customer_ids" binding:"required"`
	Message *string	`json:"message" binding:"required"`
	SubscriberID *string `json:"subscriber_id" binding:"required"`
}
