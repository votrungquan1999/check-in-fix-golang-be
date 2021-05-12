package tickets

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/requests"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
	"fmt"
)

func CreateTickets(payload requests.CreateTicketRequest) (*models.Tickets, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	newTicket := models.Tickets{
		CustomerID:  payload.CustomerID,
		ServiceID:   payload.ServiceID,
		Description: payload.Description,
		PhoneType:   payload.PhoneType,
	}

	ticketRef := firestoreClient.Collection(constants.FirestoreTicketDoc).NewDoc()

	_, err := ticketRef.Set(ctx, newTicket)
	if err != nil {
		fmt.Println(err)

		return nil, utils.ErrorInternal.New("")
	}

	createdTicketSnapshot, err := ticketRef.Get(ctx)
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	var createdTicket models.Tickets
	if err := createdTicketSnapshot.DataTo(&createdTicket); err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	return &createdTicket, nil
}
