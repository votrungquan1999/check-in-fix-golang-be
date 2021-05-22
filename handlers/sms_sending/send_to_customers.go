package sms_sending

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/requests"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
)

func SendSMSToCustomers(payload requests.SMSSendingRequest) error {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	subscriber, err := utils.GetSubscriberByID(*payload.SubscriberID)
	if err != nil {
		return err
	}

	customers := make([]models.Customers, 0)
	for _, id := range payload.CustomerIds {
		customerRef := firestoreClient.Collection(constants.FirestoreCustomerDoc).Doc(id)
		customerSnapshot, err := customerRef.Get(ctx)
		if err != nil {
			return utils.ErrorInternal.New(err.Error())
		}

		var customer models.Customers

		err = customerSnapshot.DataTo(&customer)
		if err != nil {
			return utils.ErrorInternal.New(err.Error())
		}

		customer.ID = &id

		customers = append(customers, customer)
	}

	for _, customer := range customers {
		_, err := utils.SendSMS(*subscriber.Name, *customer.PhoneNumber, *payload.Message)
		if err != nil {
			return err
		}
	}

	return nil
}
