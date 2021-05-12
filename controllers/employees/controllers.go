package employees

import (
	employeeHandler "checkinfix.com/handlers/employees"
	"checkinfix.com/requests"
	"checkinfix.com/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoutesGroup(rg *gin.RouterGroup) {
	router := rg.Group("/subscribers/:subscriber_id/employees")
	{
		router.POST("/", createEmployee)
		router.GET("/:employee_id", getEmployee)
	}
}

func createEmployee(c *gin.Context) {
	subscriberID := c.Param("subscriber_id")

	var createEmployeeRequest requests.CreateEmployeeRequest
	err := c.ShouldBindJSON(&createEmployeeRequest)
	if err != nil {
		_ = c.Error(utils.ErrorBadRequest.New(err.Error()))
		//c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	subscriber, err := employeeHandler.CreateEmployee(subscriberID, createEmployeeRequest)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": subscriber,
	})
}

func getEmployee(c *gin.Context) {
	subscriberID := c.Param("subscriber_id")
	employeeID := c.Param("employee_id")

	employee, err := employeeHandler.GetEmployee(employeeID, subscriberID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": employee,
	})
}
