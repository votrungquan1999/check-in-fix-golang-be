package subscribers

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/requests"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
)

func CreateSubscribers(payload *requests.CreateSubscriberRequest) (*models.Subscribers, error) {
	newSubscriber := models.Subscribers{
		Name:  payload.Name,
		Email: payload.Email,
	}

	ctx := context.Background()
	firestoreClient := setup.FirestoreClient
	ref := firestoreClient.Collection(constants.FirestoreSubscriberDoc).NewDoc()

	_, err := ref.Set(ctx, newSubscriber)
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	createdDoc, err := ref.Get(ctx)
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	var createdSubscriber models.Subscribers
	err = createdDoc.DataTo(&createdSubscriber)
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	createdSubscriber.ID = &createdDoc.Ref.ID

	return &createdSubscriber, nil
}
