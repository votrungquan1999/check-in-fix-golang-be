package public

import (
	"checkinfix.com/handlers/reviews"
	"checkinfix.com/handlers/settings"
	"checkinfix.com/requests"
	"checkinfix.com/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoutesGroup(rg *gin.RouterGroup) {
	reviewRouter := rg.Group("/reviews")
	{
		reviewRouter.GET("/:review_id", getReview)
		reviewRouter.POST("/:review_id/ratings", rateReview)
		reviewRouter.POST("/:review_id/feedbacks", feedbackReview)
		reviewRouter.GET("/:review_id/rating-platforms", getRatingPlatform)
	}
}

func getReview(c *gin.Context) {
	reviewID := c.Param("review_id")

	review, err := reviews.GetReviewByID(reviewID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": review,
	})
}

func rateReview(c *gin.Context) {
	reviewID := c.Param("review_id")

	var rateReviewPayload requests.RateReviewRequest
	if err := c.ShouldBindJSON(&rateReviewPayload); err != nil {
		_ = c.Error(utils.ErrorBadRequest.New(err.Error()))
		return
	}

	review, err := reviews.RateReview(reviewID, rateReviewPayload)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": review,
	})
}

func feedbackReview(c *gin.Context) {
	reviewID := c.Param("review_id")

	var feedbackReviewPayload requests.FeedbackReviewRequest
	if err := c.ShouldBindJSON(&feedbackReviewPayload); err != nil {
		_ = c.Error(utils.ErrorBadRequest.New(err.Error()))
		return
	}

	review, err := reviews.FeedbackReview(reviewID, feedbackReviewPayload)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": review,
	})
}

func getRatingPlatform(c *gin.Context) {
	reviewID := c.Param("review_id")

	ratingPlatforms, err := settings.GetRatingPlatformByReviewID(reviewID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ratingPlatforms,
	})
}
