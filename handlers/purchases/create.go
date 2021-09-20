package purchases

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/requests"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
)

func CreatePurchase(payload requests.CreatePurchaseRequest) (*models.Purchases, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	purchaseRef := firestoreClient.Collection(constants.FirestorePurchaseDoc).NewDoc()

	grandTotal, totalDiscount := CalculateCosts(payload.PurchaseProducts, payload.Discount, payload.ShippingFee)

	status := *payload.Status
	if payload.Status == nil {
		status = "Ordered"
	}

	newPurchase := models.Purchases{
		ID:               &purchaseRef.ID,
		SubscriberID:     payload.SubscriberID,
		Date:             payload.Date,
		ReferenceNumber:  payload.ReferenceNumber,
		Supplier:         payload.Supplier,
		Notes:            payload.Notes,
		Status:           &status,
		Discount:         payload.Discount,
		InstoreCredit:    payload.InstoreCredit,
		ShippingFee:      payload.ShippingFee,
		GrandTotal:       &grandTotal,
		TotalDiscount:    &totalDiscount,
		ProductPurchases: payload.PurchaseProducts,
	}

	_, err := purchaseRef.Set(ctx, newPurchase)
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	return &newPurchase, nil
}
