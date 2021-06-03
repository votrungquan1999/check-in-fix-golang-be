package customers

import (
	customerHandler "checkinfix.com/handlers/customers"
	"checkinfix.com/models"
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

	customerRouter := rg.Group("/customers")
	{
		//customerRouter.GET("/:phone_number", getCustomersByPhoneNumber)
		customerRouter.POST("", createCustomer)
		customerRouter.GET("", getCustomersHandler)
	}
}

func createCustomer(c *gin.Context) {
	//subscriberID := c.Param("subscriber_id")
	//subscriberID := c.Query("phone_number")

	var createCustomerPayload requests.CreateCustomerRequest
	if err := c.ShouldBindJSON(&createCustomerPayload); err != nil {
		_ = c.Error(utils.ErrorBadRequest.New(err.Error()))
		return
	}

	createdCustomer, err := customerHandler.CreateCustomer(createCustomerPayload)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": createdCustomer,
	})
}

func getCustomersHandler(c *gin.Context) {
	customers := getCustomers(c)
	if customers == nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": customers,
	})
}

func getCustomers(c *gin.Context) []models.Customers {
	subscriberID := c.Query("subscriber_id")

	query := c.Request.URL.Query()
	customerIDs := query["customer_id"]

	var customers []models.Customers
	var err error

	if subscriberID != "" {
		customers, err = customerHandler.GetCustomersWithSubscriberID(c, subscriberID)
		if err != nil {
			_ = c.Error(err)
			return nil
		}
		return customers
	}

	if len(customerIDs) > 0 {
		customers, err = customerHandler.GetCustomerByIDs(customerIDs)
		if err != nil {
			_ = c.Error(err)
			return nil
		}
		return customers
	}

	_ = c.Error(utils.ErrorBadRequest.New("filter is not supported"))

	return nil
}
