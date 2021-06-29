package customers

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/requests"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func BulkDeleteCustomers(payload requests.BulkDeleteCustomersRequest) ([]*models.Customers, error) {
	deletedCustomers := make([]*models.Customers, 0)
	fmt.Println(payload.Customers)

	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	err := firestoreClient.RunTransaction(ctx, func(ctx context.Context, transaction *firestore.Transaction) error {
		errorChan := make(chan error)
		deletedCustomersChan := make(chan *models.Customers)

		for _, customer := range payload.Customers {
			go func(customer requests.SingleDeleteCustomerRequest) {
				deletedCustomer, err := deleteCustomerWithTransaction(ctx, transaction, *customer.ID)
				if err != nil {
					errorChan <- err
					return
				}

				deletedCustomersChan <- deletedCustomer
			}(customer)
		}

		for {
			select {
			case err := <-errorChan:
				return err
			case deletedCustomer := <-deletedCustomersChan:
				deletedCustomers = append(deletedCustomers, deletedCustomer)
				if len(deletedCustomers) == len(payload.Customers) {
					return nil
				}
			case <-time.After(120 * time.Second):
				return utils.ErrorInternal.New("request takes too long to delete customers")
			}
		}
	})

	if err != nil {
		return nil, err
	}

	return deletedCustomers, nil
}

func deleteCustomerWithTransaction(ctx context.Context, transaction *firestore.Transaction,
	customerID string) (*models.Customers, error) {
	firestoreClient := setup.FirestoreClient

	customerRef := firestoreClient.Collection(constants.FirestoreCustomerDoc).Doc(customerID)
	customerSnapshot, err := transaction.Get(customerRef)
	if status.Code(err) == codes.NotFound {
		return nil, utils.ErrorEntityNotFound.New("customer_id not found")
	}
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	var customer models.Customers
	if err = customerSnapshot.DataTo(&customer); err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	if err = transaction.Delete(customerRef); err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	return &customer, nil
}
