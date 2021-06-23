package settings

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetRatingPlatformByReviewID(reviewID string) ([]models.RatingPlatform, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	reviewSnapshot, err := firestoreClient.Collection(constants.FirestoreReviewDoc).Doc(reviewID).Get(ctx)
	if status.Code(err) == codes.NotFound {
		return nil, utils.ErrorEntityNotFound.New("review not found")
	}
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	var review models.Reviews
	if err = reviewSnapshot.DataTo(&review); err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	ratingPlatformIter := firestoreClient.Collection(constants.FirestoreRatingPlatformDoc).Where("subscriber_id",
		"==", *review.SubscriberID).Documents(ctx)

	ratingPlatforms := make([]models.RatingPlatform, 0)
	for {
		var ratingPlatform models.RatingPlatform

		id, err := utils.GetNextDoc(ratingPlatformIter, &ratingPlatform)
		if id == "" {
			break
		}
		if err != nil {
			return nil, err
		}

		ratingPlatforms = append(ratingPlatforms, ratingPlatform)
	}

	return ratingPlatforms, nil
}
