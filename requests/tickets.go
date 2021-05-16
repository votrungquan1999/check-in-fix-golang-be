package requests

type CreateTicketRequest struct {
	CustomerID            *string `json:"customer_id" binding:"required"`
	ServiceID             *string `json:"service_id" binding:"required"`
	Description           *string `json:"description"`
	PhoneType             *string `json:"phone_type"`
	ContactPhoneNumber    *string `json:"contact_phone_number" binding:"required"`
	SMSNotificationEnable *bool   `json:"sms_notification_enable" binding:"required"`
}
