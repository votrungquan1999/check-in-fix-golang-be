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

func RateReview(ID string, ratingPayload requests.RateReviewRequest) (*models.Reviews, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	reviewRef := firestoreClient.Collection(constants.FirestoreReviewDoc).Doc(ID)
	var review models.Reviews
	if err := utils.GetDataByRef(reviewRef, &review); err != nil {
		return nil, err
	}

	if review.Rating != nil {
		return nil, utils.ErrorBadRequest.New("review is already rated")
	}

	newReviewData := models.Reviews{
		Rating:    ratingPayload.Rating,
		UpdatedAt: time.Now(),
	}
	err := utils.PatchStructData(&review, newReviewData)
	if err != nil {
		return nil, err
	}

	_, err = reviewRef.Set(ctx, review)
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	return &review, nil
}
