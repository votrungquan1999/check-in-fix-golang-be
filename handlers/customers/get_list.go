package customers

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
	"google.golang.org/api/iterator"
)

func GetCustomers(phoneNumber string, subscriberID string) ([]models.Customers, error) {
	firestoreClient := setup.FirestoreClient

	ctx := context.Background()
	iter := firestoreClient.Collection(constants.FirestoreCustomerDoc).
		Where("phone_number", "==", phoneNumber).
		Where("subscriber_id", "==", subscriberID).
		Documents(ctx)

	customers := make([]models.Customers, 0)

	for {
		var user models.Customers
		id, err := utils.GetNextDoc(iter, &user)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, utils.ErrorInternal.New(err.Error())
		}

		user.ID = &id

		customers = append(customers, user)
	}

	return customers, nil
}
