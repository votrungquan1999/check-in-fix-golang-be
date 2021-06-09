package controllers

import (
	"checkinfix.com/controllers/customers"
	"checkinfix.com/controllers/employees"
	"checkinfix.com/controllers/public"
	"checkinfix.com/controllers/reviews"
	"checkinfix.com/controllers/sms_sending"
	"checkinfix.com/controllers/subscribers"
	"checkinfix.com/controllers/tickets"
	"checkinfix.com/controllers/utils"
	"checkinfix.com/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "connect successful",
		})
	})

	publicRoutesGroup := r.Group("/public")
	{
		public.RoutesGroup(publicRoutesGroup)
	}

	privateRoutesGroup := r.Group("/private", middlewares.FirebaseAuth())
	{
		customers.RoutesGroup(privateRoutesGroup)
		subscribers.RoutesGroup(privateRoutesGroup)
		employees.RoutesGroup(privateRoutesGroup)
		utils.RoutesGroup(privateRoutesGroup)
		tickets.RoutesGroup(privateRoutesGroup)
		sms_sending.RoutesGroup(privateRoutesGroup)
		reviews.RoutesGroup(privateRoutesGroup)
	}
}
