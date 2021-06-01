package tickets

import (
	ticketHandler "checkinfix.com/handlers/tickets"
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

	//getTicketRouter := rg.Group("/subscribers/:subscriber_id")
	//{
	//
	//}

	ticketRouter := rg.Group("/tickets")
	{
		ticketRouter.POST("", createTicket)
		ticketRouter.GET("/:ticket_id/approval", approveTicket)

		ticketRouter.GET("", getTickets)
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

func getTickets(c *gin.Context) {
	//subscriberID := c.Param("subscriber_id")

	var tickets []models.Tickets
	var err error

	query := c.Request.URL.Query()

	subscriberIDs := query["subscriber_id"]
	if len(subscriberIDs) > 0 {
		tickets, err = ticketHandler.GetListTicketsBySubscriberID(subscriberIDs[0], c)
	}

	customerIDs := query["customer_id"]
	if len(customerIDs) > 0 {
		tickets, err = ticketHandler.GetTicketsByCustomerID(customerIDs[0])
	}

	if tickets == nil && err == nil{
		err = utils.ErrorBadRequest.New("tickets need to be filtered")
	}

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
