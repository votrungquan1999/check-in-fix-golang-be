package models

type Subscribers struct {
	ID    *string `json:"id" binding:"omitempty"`
	Email *string `json:"email" binding:"omitempty"`
	Name  *string `json:"name" binding:"omitempty"`
}
