package tickets

import (
	ticketHandler "checkinfix.com/handlers/tickets"
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

	//getTicketRouter := rg.Group("/subscribers/:subscriber_id")
	//{
	//
	//}

	ticketRouter := rg.Group("/tickets")
	{
		ticketRouter.POST("", createTicket)
		ticketRouter.GET("/:ticket_id/approval", approveTicket)

		ticketRouter.GET("", getTicketList)
		ticketRouter.GET("/:ticket_id", getTicketDetail)
	}
}

func createTicket(c *gin.Context) {
	var createTicketRequest requests.CreateTicketRequest

	if err := c.ShouldBindJSON(&createTicketRequest); err != nil {
		_ = c.Error(utils.ErrorBadRequest.New(err.Error()))
		return
	}

	createdTicket, err := ticketHandler.CreateTickets(createTicketRequest)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": createdTicket,
	})
}

func getTicketList(c *gin.Context) {
	subscriberID := c.Query("subscriber_id")
	customerID := c.Query("customer_id")

	tickets, err := ticketHandler.GetListTicket(subscriberID, customerID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tickets,
	})
}

func approveTicket(c *gin.Context) {
	ticketID := c.Param("ticket_id")

	ticket, err := ticketHandler.ApproveTicket(ticketID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ticket,
	})
}

func getTicketDetail(c *gin.Context) {
	ticketID := c.Param("ticket_id")

	ticket, err := ticketHandler.GetTicketDetail(ticketID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ticket,
	})
}
