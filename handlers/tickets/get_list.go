package tickets

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
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
		ID, err := utils.GetNextDoc(ticketIter, &ticket)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, utils.ErrorInternal.New(err.Error())
		}

		ticket.ID = &ID

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
		ID, err := utils.GetNextDoc(ticketIter, &ticket)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, utils.ErrorInternal.New(err.Error())
		}

		ticket.ID = &ID

		tickets = append(tickets, ticket)
	}

	return tickets, nil
}
