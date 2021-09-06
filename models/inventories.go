package models

type Product struct {
	ID            *string  `json:"id" firestore:"id"`
	SubscriberID  *string  `json:"subscriber_id" firestore:"subscriber_id"`
	ProductType   *string  `json:"product_type" firestore:"product_type"`
	ProductName   *string  `json:"product_name" firestore:"product_name"`
	ProductCode   *string  `json:"product_code" firestore:"product_code"`
	AlertQuantity *int64   `json:"alert_quantity" firestore:"alert_quantity"`
	Quantity      *int64   `json:"quantity" firestore:"quantity"`
	ProductTax    *string  `json:"product_tax" firestore:"product_tax"`
	TaxMethod     *string  `json:"tax_method" firestore:"tax_method"`
	ProductUnit   *int64   `json:"product_unit" firestore:"product_unit"`
	ProductCost   *float64 `json:"product_cost" firestore:"product_cost"`
	ProductPrice  *float64 `json:"product_price" firestore:"product_price"`
	ProductDetail *string  `json:"product_detail" firestore:"product_detail"`
	Category      *string  `json:"category" firestore:"category"`
	SubCategory   *string  `json:"sub_category" firestore:"sub_category"`
	Model         *string  `json:"model" firestore:"model"`
}
