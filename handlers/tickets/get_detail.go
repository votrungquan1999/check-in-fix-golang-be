package tickets

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
)

func GetTicketDetail(ticketID string) (*models.Tickets, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	ticketRef := firestoreClient.Collection(constants.FirestoreTicketDoc).Doc(ticketID)
	var ticket models.Tickets
	ticketSnapshot, err := ticketRef.Get(ctx)
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	err = ticketSnapshot.DataTo(&ticket)
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	return &ticket, nil
}
