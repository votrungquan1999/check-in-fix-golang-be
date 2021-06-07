package requests

type (
	CreateCustomerRequest struct {
		PhoneNumber  *string `json:"phone_number" binding:"required,numeric"`
		FirstName    *string `json:"first_name" binding:"required"`
		LastName     *string `json:"last_name" binding:"required"`
		Email        *string `json:"email"`
		AddressLine1 *string `json:"address_line_1"`
		AddressLine2 *string `json:"address_line_2"`
		City         *string `json:"city"`
		State        *string `json:"state"`
		ZipCode      *string `json:"zip_code"`
		Country      *string `json:"country"`
		SubscriberID *string `json:"subscriber_id"`
	}

	UpdateCustomerRequest struct {
		PhoneNumber  *string `json:"phone_number"`
		FirstName    *string `json:"first_name"`
		LastName     *string `json:"last_name"`
		Email        *string `json:"email"`
		AddressLine1 *string `json:"address_line_1"`
		AddressLine2 *string `json:"address_line_2"`
		City         *string `json:"city"`
		State        *string `json:"state"`
		ZipCode      *string `json:"zip_code"`
		Country      *string `json:"country"`
		SubscriberID *string `json:"subscriber_id"`
	}
)
