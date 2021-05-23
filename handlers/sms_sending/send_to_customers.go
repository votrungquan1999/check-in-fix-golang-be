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

	var subscriberID *string

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

		if subscriberID == nil {
			subscriberID = customer.SubscriberID
		} else {
			if subscriberID != customer.SubscriberID {
				return utils.ErrorBadRequest.New("all customers are not from the same subscriber")
			}
		}

		customers = append(customers, customer)
	}

	for _, customer := range customers {
		_, err := utils.SendSMSWithTwilio("", *customer.PhoneNumber, *payload.Message)
		if err != nil {
			return err
		}
	}

	return nil
}
