package tickets

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/requests"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
	"fmt"
)

func CreateTickets(payload requests.CreateTicketRequest) (*models.Tickets, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	customerRef := firestoreClient.Collection(constants.FirestoreCustomerDoc).Doc(*payload.CustomerID)
	customerSnapshot, err := customerRef.Get(ctx)
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	var customer models.Customers
	if err := customerSnapshot.DataTo(&customer); err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
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

		return nil, utils.ErrorInternal.New(err.Error())
	}

	createdTicketSnapshot, err := ticketRef.Get(ctx)
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	var createdTicket models.Tickets
	if err := createdTicketSnapshot.DataTo(&createdTicket); err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	createdTicket.ID = &ticketRef.ID

	return &createdTicket, nil
}
