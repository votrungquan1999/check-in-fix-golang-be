package employees

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/requests"
	"checkinfix.com/setup"
	"checkinfix.com/utils"
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"google.golang.org/api/iterator"
)

func CreateEmployee(subscriberID string, payload requests.CreateEmployeeRequest) (*models.Employees, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	err := firestoreClient.RunTransaction(ctx, func(ctx context.Context, transaction *firestore.Transaction) error {
		toBeCreatedUser := (&auth.UserToCreate{}).Email(*payload.Email).Password(*payload.Password)
		user, err := setup.AuthClient.CreateUser(ctx, toBeCreatedUser)
		if err != nil {
			return utils.ErrorBadRequest.New(err.Error())
		}

		ref := firestoreClient.Collection(constants.FirestoreEmployeeDoc).NewDoc()

		newEmployee := models.Employees{
			UserID:       &user.UID,
			Email:        payload.Email,
			FirstName:    payload.FirstName,
			LastName:     payload.LastName,
			SubscriberID: &subscriberID,
			Scopes:       payload.Scopes,
			ID:           &ref.ID,
		}

		err = transaction.Set(ref, newEmployee)
		if err != nil {
			return utils.ErrorBadRequest.New(err.Error())
		}

		return nil
	})
	if err != nil {
		return nil, utils.ErrorBadRequest.New(err.Error())
	}

	iter := firestoreClient.Collection(constants.FirestoreEmployeeDoc).Where("email", "==",
		*payload.Email).Documents(ctx)
	var createEmployee models.Employees
	id, err := utils.GetNextDoc(iter, &createEmployee)
	if err == iterator.Done {
		return nil, utils.ErrorInternal.New("data is not inserted for some reason")
	}
	if err != nil {
		fmt.Println(err)
		return nil, utils.ErrorInternal.New(err.Error())
	}

	createEmployee.ID = &id

	return &createEmployee, nil
}
