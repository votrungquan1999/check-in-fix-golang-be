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
