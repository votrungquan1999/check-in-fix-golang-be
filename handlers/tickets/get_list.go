package tickets

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
	"github.com/gin-gonic/gin"
)

func GetListTicketsBySubscriberID(subscriberID string, c *gin.Context) ([]models.Tickets, error) {
	tickets, err := GetDraftTickets(subscriberID)
	return tickets, err
}

func GetDraftTickets(subscriberID string) ([]models.Tickets, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	ticketIter := firestoreClient.Collection(constants.FirestoreTicketDoc).
		Where("subscriber_id", "==", subscriberID).
		Where("approved_by", "==", "").
		Documents(ctx)

	draftTickets := make([]models.Tickets, 0)
	for {
		var ticket models.Tickets
		id, err := utils.GetNextDoc(ticketIter, &ticket)
		if id == "" {
			break
		}
		if err != nil {
			return nil, utils.ErrorInternal.New(err.Error())
		}

		ticket.ID = &id

		draftTickets = append(draftTickets, ticket)
	}

	return draftTickets, nil
}

func GetTicketsByCustomerID(customerID string) ([]models.Tickets, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	ticketIter := firestoreClient.Collection(constants.FirestoreTicketDoc).
		Where("customer_id", "==", customerID).
		Documents(ctx)

	tickets := make([]models.Tickets, 0)
	for {
		var ticket models.Tickets
		id, err := utils.GetNextDoc(ticketIter, &ticket)
		if id == "" {
			break
		}
		if err != nil {
			return nil, utils.ErrorInternal.New(err.Error())
		}

		ticket.ID = &id

		tickets = append(tickets, ticket)
	}

	return tickets, nil
}
