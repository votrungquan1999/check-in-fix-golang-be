package employees

import (
	"checkinfix.com/models"
)

func GetEmployee(employeeID string, subscriberID string) (*models.Employees, error) {
	//ctx := context.Background()
	//
	//firestoreClient := setup.FirestoreClient
	//
	//employeeRef := firestoreClient.Collection(constants.FirestoreEmployeeDoc).Doc(employeeID)
	//employeeSnapshot, err := employeeRef.Get(ctx)
	////if !employeeSnapshot.Exists() {
	////
	////}
	//
	//if err != nil {
	//	if status.Code(err) == codes.NotFound {
	//		return nil, utils.ErrorEntityNotFound.New("employee id not found")
	//	}
	//	return nil, utils.ErrorInternal.New(err.Error())
	//}
	//
	//var employee models.Employees
	//err = employeeSnapshot.DataTo(&employee)
	//if err != nil {
	//	return nil, utils.ErrorInternal.New(err.Error())
	//}
	//
	//if employee.SubscriberID != &subscriberID {
	//	return nil, utils.ErrorBadRequest.New("employee is not belongs to this subscribe")
	//}
	//
	//employee.ID =

	return nil, nil
}