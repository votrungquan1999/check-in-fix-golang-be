package requests

type CreateEmployeeRequest struct {
	Email     *string  `json:"email" binding:"required"`
	FirstName *string  `json:"first_name" binding:"required"`
	LastName  *string  `json:"last_name" binding:"required"`
	Password  *string  `json:"password" binding:"required"`
	Scopes    []string `json:"scopes" binding:"required"`
}
