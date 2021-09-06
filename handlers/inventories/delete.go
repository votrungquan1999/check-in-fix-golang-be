package inventories

import (
	"checkinfix.com/constants"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func DeleteProduct(productID string) error {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	productRef := firestoreClient.Collection(constants.FirestoreProductDoc).Doc(productID)
	_, err := productRef.Get(ctx)
	if status.Code(err) == codes.NotFound {
		return utils.ErrorEntityNotFound.New("product id not found")
	}

	if _, err = productRef.Delete(ctx); err != nil {
		return utils.ErrorInternal.New(err.Error())
	}

	return nil
}
