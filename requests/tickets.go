package requests

type CreateTicketRequest struct {
	CustomerID            *string `json:"customer_id" binding:"required"`
	ServiceID             *string `json:"service_id" binding:"required"`
	Description           *string `json:"description"`
	DeviceModel           *string `json:"device_model"`
	ContactPhoneNumber    *string `json:"contact_phone_number" binding:"required"`
	SMSNotificationEnable *bool   `json:"sms_notification_enable" binding:"required"`
}
