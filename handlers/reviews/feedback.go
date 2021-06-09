package reviews

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/requests"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
	"fmt"
)

func FeedbackReview(ID string, payload requests.FeedbackReviewRequest) (*models.Reviews, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	reviewRef := firestoreClient.Collection(constants.FirestoreReviewDoc).Doc(ID)
	var review models.Reviews
	if err := utils.GetDataByRef(reviewRef, &review); err != nil {
		return nil, err
	}

	if review.Rating == nil {
		return nil, utils.ErrorBadRequest.New("review must be rated before feedback")
	}

	if *(review.Rating) > 4.0 {
		return nil, utils.ErrorBadRequest.New("review rate must be as most 4.0 to be rated")
	}

	newReviewData := models.Reviews{
		Feedback:   payload.Feedback,
		IsReviewed: true,
	}

	err := utils.PatchStructData(&review, newReviewData)
	if err != nil {
		return nil, err
	}

	fmt.Println(review)

	_, err = reviewRef.Set(ctx, review)
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	return &review, nil
}
