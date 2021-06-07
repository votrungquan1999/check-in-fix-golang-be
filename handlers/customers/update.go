package customers

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/requests"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"context"
	"fmt"
	"reflect"
)

func UpdateCustomer(customerID string, payload requests.UpdateCustomerRequest) (*models.Customers, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	customerRef := firestoreClient.Collection(constants.FirestoreCustomerDoc).Doc(customerID)

	customerSnapshot, err := customerRef.Get(ctx)
	if err != nil {
		return nil, utils.ErrorEntityNotFound.New("customer id is not found")
	}

	var customer models.Customers
	err = customerSnapshot.DataTo(&customer)
	if err != nil {
		return nil, utils.ErrorInternal.New(err.Error())
	}

	err = utils.PatchStructData(customerRef, &customer, payload)
	if err != nil {
		return nil, err
	}

	//_, err = customerRef.Set(ctx, payload)
	//if err != nil {
	//	return nil, utils.ErrorInternal.New(err.Error())
	//}

	//getExpectedCustomer(customer, payload)

	return &customer, nil
}

func getExpectedCustomer(customer models.Customers, updatePayload requests.UpdateCustomerRequest) (models.Customers,
	error) {
	t := reflect.ValueOf(updatePayload)

	fmt.Println(*t.FieldByName("PhoneNumber").Interface().(*string))
	//fmt.Println(t.FieldByName("PhoneNumber").())

	return customer, nil
}
