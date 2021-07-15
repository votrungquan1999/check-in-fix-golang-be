package subscribers

import (
	"cloud.google.com/go/firestore"
	"context"

	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/requests"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
)

func CreateSubscribers(payload *requests.CreateSubscriberRequest) (*models.Subscribers, error) {
	ctx := context.Background()
	firestoreClient := setup.FirestoreClient

	ref := firestoreClient.Collection(constants.FirestoreSubscriberDoc).NewDoc()
	newSubscriber := models.Subscribers{
		Name:  payload.Name,
		Email: payload.Email,
		ID:    &ref.ID,
	}

	err := firestoreClient.RunTransaction(ctx, func(ctx context.Context, transaction *firestore.Transaction) error {
		err := transaction.Set(ref, newSubscriber)
		if err != nil {
			return utils.ErrorInternal.New(err.Error())
		}

		err = CreateDefaultTicketStatuses(&ref.ID)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &newSubscriber, nil
}

func CreateDefaultTicketStatuses(subscriberID *string) error {
	ctx := context.Background()
	firestoreClient := setup.FirestoreClient

	err := firestoreClient.RunTransaction(ctx, func(ctx context.Context, transaction *firestore.Transaction) error {
		for _, status := range constants.DefaultTicketStatuses {
			newRef := firestoreClient.Collection(constants.FirestoreTicketStatusDoc).NewDoc()

			newTicketStatus := models.TicketStatuses{
				ID:           &newRef.ID,
				Name:         &status.Name,
				SubscriberID: subscriberID,
				Order:        &status.Order,
				Type:         utils.Int64Pointer(status.Type.Int64()),
			}

			err := transaction.Set(newRef, newTicketStatus)
			if err != nil {
				return utils.ErrorInternal.New(err.Error())
			}
		}

		return nil
	})

	return err
}
