package public

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/requests"
	"checkinfix.com/setup"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func CreateTicket(c *gin.Context) {
	var createTicketPayload requests.CreateTicketRequest
	if err := c.ShouldBindJSON(&createTicketPayload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx := context.Background()
	firestoreClient := setup.FirestoreClient

	newRef := firestoreClient.Collection(constants.FirestoreDraftTicketDoc).NewDoc()
	newTicket := models.DraftTickets{
		ID:          newRef.ID,
		UserID:      createTicketPayload.UserID,
		ServiceID:   createTicketPayload.ServiceID,
		Description: createTicketPayload.Description,
	}

	_, err := newRef.Create(ctx, newTicket)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "error save new ticket",
		})
		return
	}

	var createdTicket models.DraftTickets

	createTicketSnapShot, err := newRef.Get(ctx)
	if status.Code(err) == codes.NotFound {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "data is not created for some reason",
		})
		return
	}
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "error get created ticket",
		})
		return
	}

	err = createTicketSnapShot.DataTo(&createdTicket)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "error convert data to return to user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": createdTicket,
	})
}
