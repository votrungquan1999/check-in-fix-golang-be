package sms_sending

import (
	"checkinfix.com/handlers/sms_sending"
	"checkinfix.com/requests"
	"checkinfix.com/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoutesGroup(rg *gin.RouterGroup) {
	router := rg.Group("/sms_sending")
	{
		router.POST("/test", sendSMSTest)
		router.POST("", sendSMS)
	}
}

func sendSMSTest(c *gin.Context) {
	var sendingMessageRequest requests.SMSSendingRequest
	if err := c.ShouldBindJSON(&sendingMessageRequest); err != nil {
		_ = c.Error(utils.ErrorBadRequest.New(err.Error()))
		return
	}

	resp, err := utils.SendSMS("\nVonageAPIs", "84813792279", *sendingMessageRequest.Message)
	if err != nil {
		_ = c.Error(err)
		return
	}

	fmt.Println(resp)
}

func sendSMS(c *gin.Context) {
	var sendingMessageRequest requests.SMSSendingRequest
	if err := c.ShouldBindJSON(&sendingMessageRequest); err != nil {
		_ = c.Error(utils.ErrorBadRequest.New(err.Error()))
		return
	}

	err := sms_sending.SendSMSToCustomers(sendingMessageRequest)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
