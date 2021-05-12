package utils

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
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
