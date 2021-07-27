package requests

type CreateTicketRequest struct {
	CustomerID            *string  `json:"customer_id" binding:"required"`
	ServiceID             *string  `json:"service_id"`
	Service               *string  `json:"service"`
	Description           *string  `json:"description"`
	DeviceModel           *string  `json:"device_model"`
	ContactPhoneNumber    *string  `json:"contact_phone_number"`
	SMSNotificationEnable bool     `json:"sms_notification_enable"`
	IsDevicePowerOn       bool     `json:"is_device_power_on"`
	DroppedOffAt          *string  `json:"dropped_off_at"`
	PickUpAt              *string  `json:"pick_up_at"`
	IMEI                  *string  `json:"imei"`
	Quote                 *float64 `json:"quote"`
	Paid                  *float64 `json:"paid"`
}
