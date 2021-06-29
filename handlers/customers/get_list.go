package customers

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reflect"
	"strings"
	"time"
)

func GetCustomersWithSubscriberID(c *gin.Context, subscriberID string) ([]models.Customers, error) {
	var customers []models.Customers
	var err error

	phoneNumber := c.Query("phone_number")
	if phoneNumber != "" {
		customers, err = GetCustomersByPhoneNumber(phoneNumber, subscriberID)
		if err != nil {
			return nil, err
		}
		return customers, nil
	}

	filter := c.Query("filter")
	if filter != "" {
		customers, err = GetCustomersByFilter(strings.ToLower(filter), subscriberID)
		if err != nil {
			return nil, err
		}
		return customers, nil
	}

	customers, err = GetCustomersBySubscriberID(subscriberID)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func GetCustomersByFilter(filter string, subscriberID string) ([]models.Customers, error) {
	firestoreClient := setup.FirestoreClient

	ctx := context.Background()
	iter := firestoreClient.Collection(constants.FirestoreCustomerDoc).
		Where("subscriber_id", "==", subscriberID).
		Documents(ctx)

	customers := make([]models.Customers, 0)

	for {
		var customer models.Customers
		id, err := utils.GetNextDoc(iter, &customer)
		if id == "" {
			break
		}
		if err != nil {
			return nil, utils.ErrorInternal.New(err.Error())
		}

		if !CustomerContains(customer, filter) {
			continue
		}

		customer.ID = &id
		customers = append(customers, customer)
	}

	return customers, nil
}

func CustomerContains(customer models.Customers, key string) bool {
	v := reflect.ValueOf(customer)

	values := make([]string, 0)

	for i := 0; i < v.NumField(); i++ {
		value, ok := (v.Field(i).Interface()).(*string)

		if !ok {
			continue
		}

		if value != nil {
			values = append(values, strings.ToLower(*value))
		}
	}

	for _, value := range values {
		if strings.Contains(value, key) {
			return true
		}
	}

	return false
}

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
		if id == "" {
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

func GetCustomersBySubscriberID(subscriberID string) ([]models.Customers, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	iter := firestoreClient.Collection(constants.FirestoreCustomerDoc).
		Where("subscriber_id", "==", subscriberID).
		Documents(ctx)

	customers := make([]models.Customers, 0)

	for {
		var customer models.Customers

		id, err := utils.GetNextDoc(iter, &customer)
		if id == "" {
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
	fmt.Println(customerIDs)
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	iter := firestoreClient.Collection(constants.FirestoreCustomerDoc).
		Where("id", "in", customerIDs).
		Documents(ctx)

	customers := make([]models.Customers, 0)

	for {
		var customer models.Customers

		id, err := utils.GetNextDoc(iter, &customer)
		if id == "" {
			break
		}
		if err != nil {
			return nil, utils.ErrorInternal.New(err.Error())
		}

		customer.ID = &id

		customers = append(customers, customer)
	}

	fmt.Println(customers)

	return customers, nil
}

func GetCustomerDetailByID(customerID string) (*models.Customers, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	customerRef := firestoreClient.Collection(constants.FirestoreCustomerDoc).Doc(customerID)

	customerSnapshot, err := customerRef.Get(ctx)
	if status.Code(err) == codes.NotFound {
		return nil, utils.ErrorEntityNotFound.New("customer id not found")
	}
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	var customer models.Customers
	if err = customerSnapshot.DataTo(&customer); err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	return &customer, nil
}
