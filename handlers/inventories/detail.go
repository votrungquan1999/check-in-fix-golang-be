package inventories

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetProductDetail(productID string) (*models.Product, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	productRef := firestoreClient.Collection(constants.FirestoreProductDoc).Doc(productID)
	productSnap, err := productRef.Get(ctx)
	if status.Code(err) == codes.NotFound {
		return nil, utils.ErrorEntityNotFound.New("product id not found")
	}
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	var product models.Product
	if err = productSnap.DataTo(&product); err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	return &product, nil
}
