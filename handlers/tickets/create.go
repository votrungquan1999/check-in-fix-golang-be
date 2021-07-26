package tickets

import (
	"checkinfix.com/constants"
	"checkinfix.com/handlers/customers"
	"checkinfix.com/handlers/subscribers"
	"checkinfix.com/models"
	"checkinfix.com/requests"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"cloud.google.com/go/firestore"
	"context"
	"time"
)

func CreateTickets(payload requests.CreateTicketRequest) (*models.Tickets, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()
	var createdTicket models.Tickets

	customerRef := firestoreClient.Collection(constants.FirestoreCustomerDoc).Doc(*payload.CustomerID)
	customer, err := customers.GetCustomerDetailByID(*payload.CustomerID)
	if err != nil {
		return nil, err
	}

	ticketStatuses, err := subscribers.GetTicketStatusBySubscriberID(*customer.SubscriberID)
	if err != nil {
		return nil, err
	}

	var firstStatus = ticketStatuses[0]

	err = firestoreClient.RunTransaction(ctx, func(ctx context.Context, transaction *firestore.Transaction) error {
		if payload.ContactPhoneNumber != nil {
			customer.ContactPhoneNumber = payload.ContactPhoneNumber
			err = transaction.Set(customerRef, customer) //customerRef.Set(ctx, customer)

			if err != nil {
				return utils.ErrorInternal.New(err.Error())
			}
		}

		ticketRef := firestoreClient.Collection(constants.FirestoreTicketDoc).NewDoc()

		newTicket := prepareNewTicket(ticketRef.ID, payload, customer, firstStatus)

		err = transaction.Set(ticketRef, newTicket)
		if err != nil {
			return utils.ErrorInternal.New(err.Error())
		}

		createdTicket = newTicket

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &createdTicket, nil
}

func prepareNewTicket(id string, payload requests.CreateTicketRequest, customer *models.Customers,
	firstStatus *models.TicketStatuses) models.Tickets {
	var approvedBy = ""

	serviceID := payload.ServiceID
	if payload.ServiceID == nil {
		serviceID = utils.StringPointer(constants.ServiceDefault)
	}

	return models.Tickets{
		ID:                    &id,
		CustomerID:            payload.CustomerID,
		SubscriberID:          customer.SubscriberID,
		ServiceID:             serviceID,
		Service:               payload.Service,
		IMEI:                  payload.IMEI,
		DeviceModel:           payload.DeviceModel,
		Description:           payload.Description,
		ContactPhoneNumber:    payload.ContactPhoneNumber,
		SMSNotificationEnable: &payload.SMSNotificationEnable,
		Status:                firstStatus.Order,
		PaymentStatus:         utils.Int64Pointer(constants.TicketUnpaid),
		Quote:                 payload.Quote,
		Paid:                  payload.Paid,
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
		ApprovedBy:            &approvedBy,
		TechnicianNotes:       nil,
		Problem:               nil,
	}
}
