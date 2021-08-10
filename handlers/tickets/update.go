package tickets

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/requests"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Update(updateRequest requests.UpdateTicketRequest, ticketID string) (*models.Tickets, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	ticketRef := firestoreClient.Collection(constants.FirestoreTicketDoc).Doc(ticketID)
	ticketSnap, err := ticketRef.Get(ctx)
	if status.Code(err) == codes.NotFound {
		return nil, utils.ErrorEntityNotFound.New("ticket id not found")
	}
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	var ticket models.Tickets
	if err = ticketSnap.DataTo(&ticket); err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	updateTicket(&ticket, updateRequest)
	if _, err = ticketRef.Set(ctx, ticket); err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	return &ticket, nil
}

func updateTicket(ticket *models.Tickets, updateRequest requests.UpdateTicketRequest) {
	if updateRequest.SMSNotificationEnable != nil {
		ticket.SMSNotificationEnable = updateRequest.SMSNotificationEnable
	}
	if updateRequest.ContactPhoneNumber != nil {
		ticket.ContactPhoneNumber = updateRequest.ContactPhoneNumber
	}
	if updateRequest.Description != nil {
		ticket.Description = updateRequest.Description
	}
	if updateRequest.Devices != nil {
		devices := make([]models.TicketDevice, 0)
		for _, deviceInput := range updateRequest.Devices {
			devices = append(devices, models.TicketDevice{
				IMEI:            deviceInput.IMEI,
				DeviceModel:     deviceInput.DeviceModel,
				Service:         deviceInput.Service,
				IsDevicePowerOn: &deviceInput.IsDevicePowerOn,
			})
		}
		ticket.Devices = devices
	}
	if updateRequest.Quote != nil {
		ticket.Quote = updateRequest.Quote
	}
	if updateRequest.Paid != nil {
		ticket.Paid = updateRequest.Paid
	}
	if updateRequest.DroppedOffAt != nil {
		ticket.DroppedOffAt = updateRequest.DroppedOffAt
	}
	if updateRequest.PickUpAt != nil {
		ticket.PickUpAt = updateRequest.PickUpAt
	}
}
