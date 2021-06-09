package tickets

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/requests"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
)

func CreateTickets(payload requests.CreateTicketRequest) (*models.Tickets, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()
	var createdTicket models.Tickets

	err := firestoreClient.RunTransaction(ctx, func(ctx context.Context, transaction *firestore.Transaction) error {
		customerRef := firestoreClient.Collection(constants.FirestoreCustomerDoc).Doc(*payload.CustomerID)
		customerSnapshot, err := customerRef.Get(ctx)
		if err != nil {
			return utils.ErrorInternal.New(err.Error())
		}

		var customer models.Customers
		if err := customerSnapshot.DataTo(&customer); err != nil {
			return utils.ErrorInternal.New(err.Error())
		}

		if payload.ContactPhoneNumber != nil {
			customer.ContactPhoneNumber = payload.ContactPhoneNumber
			_, err = customerRef.Set(ctx, customer)
		}

		var approvedBy = ""

		ticketRef := firestoreClient.Collection(constants.FirestoreTicketDoc).NewDoc()

		newTicket := models.Tickets{
			ApprovedBy:            &approvedBy,
			CustomerID:            payload.CustomerID,
			ServiceID:             payload.ServiceID,
			SMSNotificationEnable: payload.SMSNotificationEnable,
			Description:           payload.Description,
			DeviceModel:           payload.DeviceModel,
			ContactPhoneNumber:    payload.ContactPhoneNumber,
			SubscriberID:          customer.SubscriberID,
			ID:                    &ticketRef.ID,
		}

		_, err = ticketRef.Set(ctx, newTicket)
		if err != nil {
			fmt.Println(err)

			return utils.ErrorInternal.New(err.Error())
		}

		createdTicketSnapshot, err := ticketRef.Get(ctx)
		if err != nil {
			return utils.ErrorInternal.New(err.Error())
		}

		if err := createdTicketSnapshot.DataTo(&createdTicket); err != nil {
			return utils.ErrorInternal.New(err.Error())
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &createdTicket, nil
}
