package requests

import "checkinfix.com/models"

type CreatePurchaseRequest struct {
	SubscriberID    *string `json:"subscriber_id"  binding:"required"`
	Date            *string `json:"date"`
	ReferenceNumber *string `json:"reference_number"`
	Supplier        *string `json:"supplier"`
	Status          *string `json:"status"  binding:"required"`

	PurchaseProducts []models.PurchaseProducts `json:"purchase_products"`

	Discount      float64 `json:"discount"`
	InstoreCredit float64 `json:"instore_credit"`
	ShippingFee   float64 `json:"shipping_fee"`
	Notes         *string `json:"notes"`
}
