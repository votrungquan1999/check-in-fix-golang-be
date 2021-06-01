package customers

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
	"google.golang.org/api/iterator"
	"time"
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

func GetCustomerByIDs(customerIDs []string) ([]models.Customers, error) {

	customers := make([]models.Customers, 0)
	customerChannel := make(chan []models.Customers)
	errorChannel := make(chan error)
	totalChunk := 0

	currentIndex := 0
	for {
		nextIndex := currentIndex + 10
		if nextIndex > len(customerIDs) {
			nextIndex = len(customerIDs)
		}

		chunkedCustomerIDs := customerIDs[currentIndex:nextIndex]

		go func() {
			chunkedCustomers, err := GetCustomerByIDsChunk(chunkedCustomerIDs)
			if err != nil {
				errorChannel <- err
			}

			customerChannel <- chunkedCustomers
		}()

		totalChunk += 1

		if nextIndex == len(customerIDs) {
			break
		}
		currentIndex = nextIndex
	}

	count := 0

	for {
		select {
		case returnCustomers := <-customerChannel:
			customers = append(customers, returnCustomers...)
			count += 1
			if count == totalChunk {
				return customers, nil
			}
		case err := <-errorChannel:
			return nil, err
		case <-time.After(30 * time.Second):
			return nil, utils.ErrorInternal.New("request takes too long to get customers")
		}
	}
}

func GetCustomerByIDsChunk(customerIDs []string) ([]models.Customers, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	iter := firestoreClient.Collection(constants.FirestoreCustomerDoc).
		Where("ID", "in", customerIDs).
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
