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
	customerRouter := rg.Group("/customers")
	{
		customerRouter.POST("", createCustomer)
		customerRouter.GET("", getCustomersHandler)
		customerRouter.GET("/:customer_id", getCustomerDetailByID)
		customerRouter.PATCH("/:customer_id", updateCustomer)
	}
}

func AdminRoutesGroup(rg *gin.RouterGroup) {
	customerRouter := rg.Group("/customers")
	{
		customerRouter.DELETE("", bulkDeleteCustomer)
	}
}

func createCustomer(c *gin.Context) {
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

func updateCustomer(c *gin.Context) {
	customerID := c.Param("customer_id")
	var updateCustomerPayload requests.UpdateCustomerRequest

	if err := c.ShouldBindJSON(&updateCustomerPayload); err != nil {
		_ = c.Error(utils.ErrorBadRequest.New(err.Error()))
		return
	}

	updatedCustomer, err := customerHandler.UpdateCustomer(customerID, updateCustomerPayload)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": updatedCustomer,
	})
}

func bulkDeleteCustomer(c *gin.Context) {
	var bulkDeleteCustomerPayload requests.BulkDeleteCustomersRequest
	if err := c.ShouldBindJSON(&bulkDeleteCustomerPayload); err != nil {
		_ = c.Error(utils.ErrorBadRequest.New(err.Error()))
		return
	}

	deletedCustomers, err := customerHandler.BulkDeleteCustomers(bulkDeleteCustomerPayload)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": deletedCustomers,
	})
}

func getCustomerDetailByID(c *gin.Context) {
	customerID := c.Param("customer_id")

	customer, err := customerHandler.GetCustomerDetailByID(customerID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": customer,
	})
}
