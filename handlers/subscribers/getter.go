package subscribers

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
)

func GetSubscriber(subscriberID string) (*models.Subscribers, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	subscriberRef := firestoreClient.Collection(constants.FirestoreSubscriberDoc).Doc(subscriberID)
	var subscriber models.Subscribers
	subscriberSnapshot, err := subscriberRef.Get(ctx)
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	err = subscriberSnapshot.DataTo(&subscriber)
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	subscriber.ID = &subscriberID

	return &subscriber, nil
}
