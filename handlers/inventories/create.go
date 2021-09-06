package inventories

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/requests"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
)

func CreateNewProduct(req requests.CreateProductRequest) (*models.Product, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	newRef := firestoreClient.Collection(constants.FirestoreProductDoc).NewDoc()

	newProduct := models.Product{
		ID:            &newRef.ID,
		SubscriberID:  req.SubscriberID,
		ProductType:   req.ProductType,
		ProductName:   req.ProductName,
		ProductCode:   req.ProductCode,
		AlertQuantity: req.AlertQuantity,
		Quantity:      req.Quantity,
		ProductTax:    req.ProductTax,
		TaxMethod:     req.TaxMethod,
		ProductUnit:   req.ProductUnit,
		ProductCost:   req.ProductCost,
		ProductPrice:  req.ProductPrice,
		ProductDetail: req.ProductDetail,
		Category:      req.Category,
		SubCategory:   req.SubCategory,
		Model:         req.Model,
	}

	_, err := newRef.Set(ctx, newProduct)
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	return &newProduct, nil
}
