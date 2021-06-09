package reviews

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/requests"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
)

func CreateReview(payload requests.CreateReviewRequest) (*models.Reviews, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	customerRef := firestoreClient.Collection(constants.FirestoreCustomerDoc).Doc(*payload.CustomerID)
	var customer models.Customers

	err := utils.GetDataByRef(customerRef, &customer)
	if err != nil {
		return nil, err
	}

	newRef := firestoreClient.Collection(constants.FirestoreReviewDoc).NewDoc()

	newReview := models.Reviews{
		ID:           &newRef.ID,
		CustomerID:   payload.CustomerID,
		SubscriberID: customer.SubscriberID,
	}

	_, err = newRef.Set(ctx, newReview)
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	var createdReview models.Reviews
	err = utils.GetDataByRef(newRef, &createdReview)
	if err != nil {
		return nil, err
	}

	return &createdReview, nil
}
