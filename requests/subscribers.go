package requests

type CreateSubscriberRequest struct {
	Email *string `json:"email" binding:"required"`
	Name  *string `json:"name" binding:"required"`
}
