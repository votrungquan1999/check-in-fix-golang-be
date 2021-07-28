package models

import "time"

type Tickets struct {
	ID                    *string  `firestore:"id,omitempty" json:"id"`
	CustomerID            *string  `firestore:"customer_id,omitempty" json:"customer_id"`
	SubscriberID          *string  `firestore:"subscriber_id,omitempty" json:"subscriber_id"`
	ServiceID             *string  `firestore:"service_id,omitempty" json:"service_id"`
	Description           *string  `firestore:"description,omitempty" json:"description"`
	ContactPhoneNumber    *string  `firestore:"contact_phone_number,omitempty" json:"contact_phone_number"`
	SMSNotificationEnable *bool    `firestore:"sms_notification_enable,omitempty" json:"sms_notification_enable"`
	DroppedOffAt          *string  `firestore:"dropped_off_at,omitempty" json:"dropped_off_at"`
	PickUpAt              *string  `firestore:"pick_up_at" json:"pick_up_at"`
	Status                *int64   `firestore:"status" json:"status"`
	PaymentStatus         *int64   `firestore:"payment_status" json:"payment_status"`
	Quote                 *float64 `firestore:"quote" json:"quote"`
	Paid                  *float64 `firestore:"paid" json:"paid"`

	Devices []TicketDevice `firestore:"devices" json:"devices"`

	CreatedAt time.Time `firestore:"created_at" json:"created_at"`
	UpdatedAt time.Time `firestore:"updated_at" json:"updated_at"`

	ApprovedBy      *string `firestore:"approved_by,omitempty" json:"approved_by"`
	TechnicianNotes *string `firestore:"technician_notes" json:"technician_notes"`
	Problem         *string `firestore:"problem" json:"problem"`
}

type TicketDevice struct {
	IMEI            *string `firestore:"imei" json:"imei"`
	DeviceModel     *string `firestore:"device_model,omitempty" json:"device_model"`
	Service         *string `firestore:"service,omitempty" json:"service"`
	IsDevicePowerOn *bool   `firestore:"is_device_power_on,omitempty" json:"is_device_power_on"`
}
