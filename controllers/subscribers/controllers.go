package subscribers

import (
	subscriberHandler "checkinfix.com/handlers/subscribers"
	"checkinfix.com/requests"
	"checkinfix.com/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoutesGroup(rg *gin.RouterGroup) {
	router := rg.Group("/subscribers")
	{
		router.POST("/", createSubscriber)
		router.GET("/:subscriber_id", getSubscriber)
	}
}

func createSubscriber(c *gin.Context) {
	var createSubscriberRequest requests.CreateSubscriberRequest
	err := c.ShouldBindJSON(&createSubscriberRequest)
	if err != nil {
		_ = c.Error(utils.ErrorBadRequest.New(err.Error()))
		//c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	subscriber, err := subscriberHandler.CreateSubscribers(&createSubscriberRequest)
	if err != nil {
		_ = c.Error(err)
		//c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": subscriber,
	})
}

func getSubscriber(c *gin.Context) {
	subscriberID := c.Param("subscriber_id")

	subscriber, err := subscriberHandler.GetSubscriber(subscriberID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": subscriber,
	})
}
