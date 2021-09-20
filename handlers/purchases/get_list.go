package purchases

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
)

func GetListPurchases(subscriberID string) ([]*models.Purchases, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	purchaseIter := firestoreClient.Collection(constants.FirestorePurchaseDoc).Where("subscriber_id", "==",
		subscriberID).Documents(ctx)

	purchases := make([]*models.Purchases, 0)
	for {
		var product models.Purchases
		id, err := utils.GetNextDoc(purchaseIter, &product)
		if err != nil {
			return nil, err
		}
		if id == "" {
			break
		}

		purchases = append(purchases, &product)
	}

	return purchases, nil
}
