package requests

type (
	CreateCustomerRequest struct {
		PhoneNumber *string `json:"phone_number" binding:"required,numeric"`
		FirstName   *string `json:"first_name" binding:"required"`
		LastName    *string `json:"last_name" binding:"required"`
		Email       *string `json:"email" binding:"required"`
	}
)
