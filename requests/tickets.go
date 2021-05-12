package requests

type CreateTicketRequest struct {
	CustomerID  *string `json:"customer_id" binding:"required"`
	ServiceID   *string `json:"service_id" binding:"required"`
	Description *string `json:"description" binding:"required"`
	PhoneType   *string `json:"phone_type" binding:"required"`
}
