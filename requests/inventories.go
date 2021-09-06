package requests

type CreateProductRequest struct {
	SubscriberID  *string  `json:"subscriber_id" binding:"required"`
	ProductType   *string  `json:"product_type"`
	ProductName   *string  `json:"product_name" binding:"required"`
	ProductCode   *string  `json:"product_code"`
	AlertQuantity *int64   `json:"alert_quantity"`
	Quantity      *int64   `json:"quantity"`
	ProductTax    *string  `json:"product_tax"`
	TaxMethod     *string  `json:"tax_method"`
	ProductUnit   *int64   `json:"product_unit"`
	ProductCost   *float64 `json:"product_cost"`
	ProductPrice  *float64 `json:"product_price"`
	ProductDetail *string  `json:"product_detail"`
	Category      *string  `json:"category"`
	SubCategory   *string  `json:"sub_category"`
	Model         *string  `json:"model"`
}

type UpdateProductRequest struct {
	ProductType   *string  `json:"product_type"`
	ProductName   *string  `json:"product_name"`
	ProductCode   *string  `json:"product_code"`
	AlertQuantity *int64   `json:"alert_quantity"`
	Quantity      *int64   `json:"quantity"`
	ProductTax    *string  `json:"product_tax"`
	TaxMethod     *string  `json:"tax_method"`
	ProductUnit   *int64   `json:"product_unit"`
	ProductCost   *float64 `json:"product_cost"`
	ProductPrice  *float64 `json:"product_price"`
	ProductDetail *string  `json:"product_detail"`
	Category      *string  `json:"category"`
	SubCategory   *string  `json:"sub_category"`
	Model         *string  `json:"model"`
}
