package reviews

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/requests"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func BulkCreateReview(payload requests.BulkCreateReviewRequest) ([]*models.Reviews, error) {
	reviews := make([]*models.Reviews, 0)

	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	err := firestoreClient.RunTransaction(ctx, func(ctx context.Context, transaction *firestore.Transaction) error {
		reviewChan := make(chan *models.Reviews)
		errorChan := make(chan error)

		for _, singleCustomer := range payload.Customers {
			go func(item requests.CreateSingleReviewRequest) {
				createdReview, err := createReviewWithTransaction(ctx, transaction, item)
				if err != nil {
					errorChan <- err
					return
				}

				reviewChan <- createdReview
			}(singleCustomer)
		}

		total := 0

		for {
			select {
			case err := <-errorChan:
				return err
			case returnReview := <-reviewChan:
				reviews = append(reviews, returnReview)
				total += 1
				if total == len(payload.Customers) {
					return nil
				}
			case <-time.After(120 * time.Second):
				return utils.ErrorInternal.New("request takes too long to create reviews")
			}
		}
	})
	if err != nil {
		return nil, err
	}

	return reviews, nil

}

func CreateReview(payload requests.CreateSingleReviewRequest) (*models.Reviews, error) {
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
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
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

func createReviewWithTransaction(ctx context.Context, transaction *firestore.Transaction,
	payload requests.CreateSingleReviewRequest) (*models.Reviews, error) {

	firestoreClient := setup.FirestoreClient

	customerRef := firestoreClient.Collection(constants.FirestoreCustomerDoc).Doc(*payload.CustomerID)
	var customer models.Customers
	customerSnapshot, err := transaction.Get(customerRef)
	if status.Code(err) == codes.NotFound {
		return nil, utils.ErrorEntityNotFound.New("customer_id not found")
	}
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	if err = customerSnapshot.DataTo(&customer); err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	newRef := firestoreClient.Collection(constants.FirestoreReviewDoc).NewDoc()

	newReview := models.Reviews{
		ID:           &newRef.ID,
		CustomerID:   payload.CustomerID,
		SubscriberID: customer.SubscriberID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err = transaction.Set(newRef, newReview)
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	return &newReview, nil
}
