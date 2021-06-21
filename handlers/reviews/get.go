package reviews

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
	"sort"
)

func GetReviewByID(ID string) (*models.Reviews, error) {
	firestoreClient := setup.FirestoreClient

	reviewRef := firestoreClient.Collection(constants.FirestoreReviewDoc).Doc(ID)
	var review models.Reviews
	if err := utils.GetDataByRef(reviewRef, &review); err != nil {
		return nil, err
	}

	return &review, nil
}

func GetReviewList(subscriberID string, currentPage int, limit int) ([]models.Reviews, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()
	reviewCollection := firestoreClient.
		Collection(constants.FirestoreReviewDoc).Where("subscriber_id", "==", subscriberID)

	query := reviewCollection

	reviewsIter := query.Documents(ctx)

	reviewsList := make([]models.Reviews, 0)
	for {
		var review models.Reviews
		id, err := utils.GetNextDoc(reviewsIter, &review)
		if id == "" {
			break
		}
		if err != nil {
			return nil, err
		}

		reviewsList = append(reviewsList, review)
	}

	sort.Slice(reviewsList, func(i, j int) bool {
		return reviewsList[i].UpdatedAt.After(reviewsList[j].UpdatedAt)
	})

	return reviewsList, nil
}
