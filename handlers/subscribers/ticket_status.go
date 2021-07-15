package subscribers

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"cloud.google.com/go/firestore"
	"context"
)

func GetTicketStatusBySubscriberID(subscriberID string) ([]*models.TicketStatuses, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	var subscriberTicketStatusesIter = firestoreClient.Collection(constants.FirestoreTicketStatusDoc).Where(
		"subscriber_id", "==", subscriberID).OrderBy("order", firestore.Asc).Documents(ctx)

	ticketStatuses := make([]*models.TicketStatuses, 0)
	for {
		var currentTicketStatus models.TicketStatuses
		id, err := utils.GetNextDoc(subscriberTicketStatusesIter, &currentTicketStatus)
		if err != nil {
			return nil, err
		}
		if id == "" {
			break
		}

		ticketStatuses = append(ticketStatuses, &currentTicketStatus)
	}

	if len(ticketStatuses) == 0 {
		return nil, utils.ErrorBadRequest.New("please create default ticket status for subscriber")
	}

	return ticketStatuses, nil
}
