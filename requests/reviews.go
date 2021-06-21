package requests

type BulkCreateReviewRequest struct {
	Customers []*CreateSingleReviewRequest `json:"customers" binding:"required,min=1,max=50,dive"`
}

type CreateSingleReviewRequest struct {
	CustomerID *string `json:"id" binding:"required"`
}

type RateReviewRequest struct {
	Rating *float64 `json:"rating" binding:"required"`
}

type FeedbackReviewRequest struct {
	Feedback *string `json:"feedback" binding:"required"`
}
