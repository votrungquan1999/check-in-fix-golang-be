package tickets

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
	"time"
)

func GetListTicket(subscriberID string, customerID string) ([]*models.Tickets, error) {
	if subscriberID != "" {
		tickets, err := getListTicketsBySubscriberID(subscriberID)
		return tickets, err
	}

	if customerID != "" {
		tickets, err := getTicketsByCustomerID(customerID)
		return tickets, err
	}

	return nil, utils.ErrorBadRequest.New("tickets need to be filtered")
}

func getListTicketsBySubscriberID(subscriberID string) ([]*models.Tickets, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	ticketIter := firestoreClient.Collection(constants.FirestoreTicketDoc).
		Where("subscriber_id", "==", subscriberID).
		Where("approved_by", "==", "").
		Documents(ctx)

	tickets := make([]*models.Tickets, 0)
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

		tickets = append(tickets, &ticket)
	}

	return tickets, nil
}

func getTicketsByCustomerID(customerID string) ([]*models.Tickets, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	ticketIter := firestoreClient.Collection(constants.FirestoreTicketDoc).
		Where("customer_id", "==", customerID).
		Where("to_be_removed_at", "!=", time.Now()).
		Documents(ctx)

	tickets := make([]*models.Tickets, 0)
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

		tickets = append(tickets, &ticket)
	}

	return tickets, nil
}
