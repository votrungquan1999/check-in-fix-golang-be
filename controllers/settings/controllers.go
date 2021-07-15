package settings

import (
	"checkinfix.com/handlers/settings"
	"checkinfix.com/handlers/subscribers"
	"checkinfix.com/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoutesGroup(rg *gin.RouterGroup) {
	router := rg.Group("/settings")
	{
		router.GET("/rating-platforms", getRatingPlatform)
		router.GET("/ticket-statuses", getTicketStatuses)
	}
}

func getRatingPlatform(c *gin.Context) {
	subscriberID := c.Query("subscriber_id")
	if subscriberID == "" {
		_ = c.Error(utils.ErrorBadRequest.New("subscriber_id cannot be empty"))
		return
	}

	ratingPlatforms, err := settings.GetRatingPlatformBySubscriberID(subscriberID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ratingPlatforms,
	})
}

func getTicketStatuses(c *gin.Context) {
	subscriberID := c.Query("subscriber_id")
	if subscriberID == "" {
		_ = c.Error(utils.ErrorBadRequest.New("subscriber_id cannot be empty"))
		return
	}

	ticketStatuses, err := subscribers.GetTicketStatusBySubscriberID(subscriberID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ticketStatuses,
	})
}
