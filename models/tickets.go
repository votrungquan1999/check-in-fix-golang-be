package models

type Tickets struct {
	ID                    *string `firestore:"id,omitempty" json:"id,omitempty"`
	CustomerID            *string `firestore:"customer_id,omitempty" json:"customer_id"`
	SubscriberID          *string `firestore:"subscriber_id,omitempty" json:"subscriber_id"`
	ServiceID             *string `firestore:"service_id,omitempty" json:"service_id"`
	Description           *string `firestore:"description,omitempty" json:"description"`
	ApprovedBy            *string `firestore:"approved_by,omitempty" json:"approved_by"`
	DeviceModel           *string `firestore:"device_model,omitempty" json:"device_model"`
	ContactPhoneNumber    *string `firestore:"contact_phone_number,omitempty" json:"contact_phone_number"`
	SMSNotificationEnable *bool   `firestore:"sms_notification_enable,omitempty" json:"sms_notification_enable"`

	TechnicianNotes *string `firestore:"technician_notes" json:"technician_notes"`
	Problem         *string `firestore:"problem" json:"problem"`
}

type DraftTickets struct {
	ID          string `firestore:"id,omitempty" json:"id,omitempty"`
	UserID      string `firestore:"user_id,omitempty" json:"user_id,omitempty"`
	ServiceID   int    `firestore:"service_id,omitempty" json:"service_id,omitempty"`
	Description string `firestore:"description,omitempty" json:"description,omitempty"`
}
