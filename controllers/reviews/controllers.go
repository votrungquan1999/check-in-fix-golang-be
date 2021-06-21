package reviews

import (
	"checkinfix.com/handlers/reviews"
	"checkinfix.com/requests"
	"checkinfix.com/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoutesGroup(rg *gin.RouterGroup) {
	reviewRouter := rg.Group("/reviews")
	{
		//reviewRouter.POST("")
		//reviewRouter.GET("/:review_id", getReview)
		reviewRouter.POST("", bulkCreateReviewForCustomer)
		reviewRouter.GET("", utils.WithPagination(), getReviewList)
	}

	//customerReviewRouter := rg.Group("/customers/:customer_id/reviews")
	//{
	//}
}

func bulkCreateReviewForCustomer(c *gin.Context) {
	var bulkCreateReviewPayload requests.BulkCreateReviewRequest

	if err := c.ShouldBindJSON(&bulkCreateReviewPayload); err != nil {
		_ = c.Error(utils.ErrorBadRequest.New(err.Error()))
		return
	}

	createdReviews, err := reviews.BulkCreateReview(bulkCreateReviewPayload)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": createdReviews,
	})
}

func getReviewList(c *gin.Context) {
	subscriberID := c.Query("subscriber_id")
	if subscriberID == "" {
		_ = c.Error(utils.ErrorBadRequest.New("missing subscriber_id query param"))
		return
	}

	reviewList, err := reviews.GetReviewList(subscriberID, c.GetInt("current_page"), c.GetInt("limit"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": reviewList,
	})
}
