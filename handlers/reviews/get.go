package reviews

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
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
