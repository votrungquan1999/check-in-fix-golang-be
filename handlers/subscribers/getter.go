package subscribers

import (
	"context"

	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetSubscriber(subscriberID string) (*models.Subscribers, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	subscriberRef := firestoreClient.Collection(constants.FirestoreSubscriberDoc).Doc(subscriberID)
	var subscriber models.Subscribers
	subscriberSnapshot, err := subscriberRef.Get(ctx)
	if status.Code(err) == codes.NotFound {
		return nil, utils.ErrorEntityNotFound.New("get subscriber not found")
	}
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
