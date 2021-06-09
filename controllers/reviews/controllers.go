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
		reviewRouter.POST("", createReviewForCustomer)
	}

	//customerReviewRouter := rg.Group("/customers/:customer_id/reviews")
	//{
	//}
}



func createReviewForCustomer(c *gin.Context) {
	var createReviewPayload requests.CreateReviewRequest

	if err := c.ShouldBindJSON(&createReviewPayload); err != nil {
		_ = c.Error(utils.ErrorBadRequest.New(err.Error()))
		return
	}

	review, err := reviews.CreateReview(createReviewPayload)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": review,
	})

}
