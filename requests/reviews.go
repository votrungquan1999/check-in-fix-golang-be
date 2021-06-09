package requests

type CreateReviewRequest struct {
	CustomerID *string `json:"customer_id" binding:"required"`
}

type RateReviewRequest struct {
	Rating *float64 `json:"rating" binding:"required"`
}

type FeedbackReviewRequest struct {
	Feedback *string `json:"feedback" binding:"required"`
}
