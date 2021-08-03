package workers

import (
	"checkinfix.com/constants"
	"checkinfix.com/setup"
	"context"
	"google.golang.org/api/iterator"
	"time"
)

func StartTicketWorkers() {
	deleteDraftTicket()
}

func deleteDraftTicket() {
	go func() {
		for {
			select {
			case <-time.After(1 * time.Hour):
				deleteDraftTicket()
			}
		}
	}()

	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	ticketIter := firestoreClient.Collection(constants.FirestoreTicketDoc).Where("to_be_removed_at", "<",
		time.Now()).Documents(ctx)

	for {
		ticketSnap, err := ticketIter.Next()
		if err == iterator.Done {
			return
		}
		if err != nil {
			return
		}

		_, err = ticketSnap.Ref.Delete(ctx)
		if err != nil {
			return
		}
	}

}
