package utils

import (
	"checkinfix.com/constants"
	"checkinfix.com/handlers/subscribers"
	"checkinfix.com/models"
	"checkinfix.com/requests"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoutesGroup(rg *gin.RouterGroup) {
	authRouter := rg.Group("/auth")
	{
		authRouter.GET("/employee-info", getEmployInfo)
	}

	utilsRouter := rg.Group("/utils")
	{
		utilsRouter.POST("/default-ticket-status", createDefaultTicketStatus)
	}
}

func getEmployInfo(c *gin.Context) {
	userInterface, ok := c.Get("user")
	if !ok {
		_ = c.Error(utils.ErrorBadRequest.New("you need to login to use this api"))
	}

	user, ok := userInterface.(*auth.UserRecord)
	if !ok {
		_ = c.Error(utils.ErrorInternal.New("can not get user from auth"))
	}

	firestoreClient := setup.FirestoreClient
	employeeIter := firestoreClient.Collection(constants.FirestoreEmployeeDoc).
		Where("email", "==", user.Email).Documents(context.Background())

	var employee models.Employees
	id, err := utils.GetNextDoc(employeeIter, &employee)
	if err != nil {
		_ = c.Error(utils.ErrorInternal.New(err.Error()))
	}

	employee.ID = &id

	c.JSON(http.StatusOK, gin.H{
		"data": employee,
	})
}

func createDefaultTicketStatus(c *gin.Context) {
	var createDefaultTicketPayload requests.CreateDefaultTickets
	if err := c.ShouldBindJSON(&createDefaultTicketPayload); err != nil {
		_ = c.Error(utils.ErrorBadRequest.New(err.Error()))
		return
	}

	if err := subscribers.CreateDefaultTicketStatuses(createDefaultTicketPayload.SubscriberID); err != nil {
		_ = c.Error(err)
		return
	}

	c.AbortWithStatus(204)
}
