package tickets

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Delete(ticketID string) (*models.Tickets, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	ticketRef := firestoreClient.Collection(constants.FirestoreTicketDoc).Doc(ticketID)
	ticketSnap, err := ticketRef.Get(ctx)
	if status.Code(err) == codes.NotFound {
		return nil, utils.ErrorEntityNotFound.New("ticket id not found")
	}
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	var ticket models.Tickets
	if err = ticketSnap.DataTo(&ticket); err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	if _, err = ticketRef.Delete(ctx); err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	return &ticket, err
}
