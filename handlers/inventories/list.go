package inventories

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
)

func GetProductBySubscriber(subscriberID string) ([]*models.Product, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	productIter := firestoreClient.Collection(constants.FirestoreProductDoc).Where("subscriber_id", "==",
		subscriberID).Documents(ctx)

	products := make([]*models.Product, 0)

	for {
		var product models.Product
		id, err := utils.GetNextDoc(productIter, &product)
		if err != nil {
			return nil, err
		}
		if id == "" {
			break
		}

		products = append(products, &product)
	}

	return products, nil
}
