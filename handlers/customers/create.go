package customers

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/requests"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
)

func CreateCustomer(payload requests.CreateCustomerRequest) (*models.Customers, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	customerExisted, err := doesCustomerExisted(payload)
	if err != nil {
		return nil, err
	}

	if customerExisted {
		return nil, utils.ErrorBadRequest.New("Customer with the same phone number and name already exists")
	}

	ref := firestoreClient.Collection(constants.FirestoreCustomerDoc).NewDoc()

	newCustomer := models.Customers{
		PhoneNumber:  payload.PhoneNumber,
		FirstName:    payload.FirstName,
		LastName:     payload.LastName,
		Email:        payload.Email,
		SubscriberID: payload.SubscriberID,
		AddressLine1: payload.AddressLine1,
		AddressLine2: payload.AddressLine2,
		City:         payload.City,
		State:        payload.State,
		ZipCode:      payload.ZipCode,
		Country:      payload.Country,
		ID:           &ref.ID,
	}

	_, err = ref.Set(ctx, newCustomer)
	if err != nil {
		fmt.Println(err)

		return nil, utils.ErrorInternal.New("")
	}

	var createdCustomer models.Customers
	createdCustomerSnapshot, err := ref.Get(ctx)
	if err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{
		//	"error": "data is not created for some reason",
		//})
		return nil, utils.ErrorInternal.New(err.Error())
	}

	err = createdCustomerSnapshot.DataTo(&createdCustomer)
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	createdCustomer.ID = &ref.ID

	return &createdCustomer, nil
}

func doesCustomerExisted(createCustomerPayload requests.CreateCustomerRequest) (bool, error) {
	//fmt.Println(subscriberID)
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()
	iter := firestoreClient.
		Collection(constants.FirestoreCustomerDoc).
		Where("phone_number", "==", createCustomerPayload.PhoneNumber).
		Where("first_name", "==", createCustomerPayload.FirstName).
		Where("last_name", "==", createCustomerPayload.LastName).
		Where("subscriber_id", "==", createCustomerPayload.SubscriberID).
		Documents(ctx)

	_, err := iter.Next()

	if err == iterator.Done {
		return false, nil
	}

	if err != nil {
		return false, utils.ErrorInternal.New(err.Error())
	}

	return true, nil
}
