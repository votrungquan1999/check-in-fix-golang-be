package models

type Purchases struct {
	ID              *string `json:"id" firestore:"id"`
	SubscriberID    *string `json:"subscriber_id" firestore:"subscriber_id"`
	Date            *string `json:"date" firestore:"date"`
	ReferenceNumber *string `json:"reference_number" firestore:"reference_number"`
	Supplier        *string `json:"supplier" firestore:"supplier"`
	Notes           *string `json:"notes" firestore:"notes"`
	Status          *string `json:"status" firestore:"status"`

	Discount      float64 `json:"discount" firestore:"discount"`
	InstoreCredit float64 `json:"instore_credit" firestore:"instore_credit"`
	ShippingFee   float64 `json:"shipping_fee" firestore:"shipping_fee"`

	GrandTotal    *float64 `json:"grand_total" firestore:"grand_total"`
	TotalDiscount *float64 `json:"total_discount" firestore:"total_discount"`

	ProductPurchases []PurchaseProducts `json:"product_purchases" firestore:"product_purchases"`
}

type PurchaseProducts struct {
	ProductID   *string `json:"product_id" firestore:"product_id"`
	ProductCode *string `json:"product_code" firestore:"product_code"`
	ProductName *string `json:"product_name" firestore:"product_name"`

	UnitCost float64 `json:"unit_cost" firestore:"unit_cost"`
	Quantity int64   `json:"quantity" firestore:"quantity"`
	Discount float64 `json:"discount" firestore:"discount"`
	TaxType  *string `json:"tax_type" firestore:"tax_type" binding:"required"`
	TaxValue float64 `json:"tax_value" firestore:"tax_value"`
}
