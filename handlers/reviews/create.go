package reviews

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/requests"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
	"time"
)

func BulkCreateReview(payload requests.BulkCreateReviewRequest) ([]*models.Reviews, error) {
	reviewChan := make(chan *models.Reviews)
	errorChan := make(chan error)

	for _, singleCustomer := range payload.Customers {
		item := singleCustomer
		go func() {
			createdReview, err := CreateReview(item)
			if err != nil {
				errorChan <- err
			}

			reviewChan <- createdReview
		}()
	}

	reviews := make([]*models.Reviews, 0)
	total := 0

	for {
		select {
		case err := <-errorChan:
			return nil, err
		case returnReview := <-reviewChan:
			reviews = append(reviews, returnReview)
			total += 1
			if total == len(payload.Customers) {
				return reviews, nil
			}
		case <-time.After(120 * time.Second):
			return nil, utils.ErrorInternal.New("request takes too long to create reviews")
		}
	}
}

func CreateReview(payload *requests.CreateSingleReviewRequest) (*models.Reviews, error) {
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
