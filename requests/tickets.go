package requests

type CreateTicketRequest struct {
	CustomerID            *string  `json:"customer_id" binding:"required"`
	ServiceID             *string  `json:"service_id"`
	Description           *string  `json:"description"`
	ContactPhoneNumber    *string  `json:"contact_phone_number"`
	SMSNotificationEnable bool     `json:"sms_notification_enable"`
	DroppedOffAt          *string  `json:"dropped_off_at"`
	PickUpAt              *string  `json:"pick_up_at"`
	Quote                 *float64 `json:"quote"`
	Paid                  *float64 `json:"paid"`

	Devices []*DeviceInput `json:"devices"`
}

type DeviceInput struct {
	IsDevicePowerOn bool    `json:"is_device_power_on"`
	IMEI            *string `json:"imei"`
	Service         *string `json:"service"`
	DeviceModel     *string `json:"device_model"`
}
