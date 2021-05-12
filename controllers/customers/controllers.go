package customers

import (
	customerHandler "checkinfix.com/handlers/customers"
	"checkinfix.com/requests"
	"checkinfix.com/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoutesGroup(rg *gin.RouterGroup) {
	//ticketRouter := rg.Group("/tickets")
	//{
	//	ticketRouter.POST("/", public.CreateTicket)
	//}

	customerRouter := rg.Group("/subscribers/:subscriber_id/customers")
	{
		customerRouter.GET("/:phone_number", getCustomers)
		customerRouter.POST("/", createCustomer)
	}
}

func createCustomer(c *gin.Context) {
	subscriberID := c.Param("subscriber_id")

	var createCustomerPayload requests.CreateCustomerRequest
	if err := c.ShouldBindJSON(&createCustomerPayload); err != nil {
		_ = c.Error(utils.ErrorBadRequest.New(err.Error()))
		return
	}

	createdCustomer, err := customerHandler.CreateCustomer(subscriberID, createCustomerPayload)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": createdCustomer,
	})
}

func getCustomers(c *gin.Context) {
	subscriberID := c.Param("subscriber_id")
	phoneNumber := c.Param("phone_number")

	customers, err := customerHandler.GetCustomers(phoneNumber, subscriberID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": customers,
	})
}
