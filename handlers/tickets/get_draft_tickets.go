package tickets

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
	"google.golang.org/api/iterator"
)

func GetDraftTickets(subscriberID string) ([]models.Tickets, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	ticketIter := firestoreClient.Collection(constants.FirestoreTicketDoc).
		//EndBefore(true).
		//OrderBy("approved_by", firestore.Asc).
		Where("subscriber_id", "==", subscriberID).
		Where("approved_by", "==", "").
		Documents(ctx)

	draftTickets := make([]models.Tickets, 0)
	for {
		var newTicket models.Tickets
		_, err := utils.GetNextDoc(ticketIter, &newTicket)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, utils.ErrorInternal.New(err.Error())
		}

		//newTicket.ID = &ID

		draftTickets = append(draftTickets, newTicket)
	}

	return draftTickets, nil
}
