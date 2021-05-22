package customers

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
	"google.golang.org/api/iterator"
)

func GetCustomersByPhoneNumber(phoneNumber string, subscriberID string) ([]models.Customers, error) {
	firestoreClient := setup.FirestoreClient

	ctx := context.Background()
	iter := firestoreClient.Collection(constants.FirestoreCustomerDoc).
		Where("phone_number", "==", phoneNumber).
		Where("subscriber_id", "==", subscriberID).
		Documents(ctx)

	customers := make([]models.Customers, 0)

	for {
		var customer models.Customers
		id, err := utils.GetNextDoc(iter, &customer)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, utils.ErrorInternal.New(err.Error())
		}

		customer.ID = &id

		customers = append(customers, customer)
	}

	return customers, nil
}

func GetCustomers(subscriberID string) ([]models.Customers, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	iter := firestoreClient.Collection(constants.FirestoreCustomerDoc).
		Where("subscriber_id", "==", subscriberID).
		Documents(ctx)

	customers := make([]models.Customers, 0)

	for {
		var customer models.Customers

		id, err := utils.GetNextDoc(iter, &customer)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, utils.ErrorInternal.New(err.Error())
		}

		customer.ID = &id

		customers = append(customers, customer)
	}

	return customers, nil
}
