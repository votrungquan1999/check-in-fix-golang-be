package utils

import (
	"checkinfix.com/constants"
	"checkinfix.com/models"
	"checkinfix.com/setup"
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reflect"
)

func GetNextDoc(iter *firestore.DocumentIterator, returnObj interface{}) (string, error) {
	doc, err := iter.Next()

	if err == iterator.Done {
		return "", err
	}

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	err = doc.DataTo(returnObj)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return doc.Ref.ID, nil
}

func GetDataByRef(ref *firestore.DocumentRef, returnObj interface{}) error {
	ctx := context.Background()
	snapShot, err := ref.Get(ctx)
	if err != nil {
		return ErrorEntityNotFound.New("id is not found")
	}

	if err = snapShot.DataTo(returnObj); err != nil {
		return ErrorInternal.New(err.Error())
	}

	return nil
}

func GetSubscriberByID(id string) (*models.Subscribers, error) {
	firestoreClient := setup.FirestoreClient
	ctx := context.Background()

	subscriberRef := firestoreClient.Collection(constants.FirestoreSubscriberDoc).Doc(id)
	subscriberSnapshot, err := subscriberRef.Get(ctx)
	if status.Code(err) == codes.NotFound {
		return nil, ErrorEntityNotFound.New("subscriber not found")
	}
	if err != nil {
		return nil, ErrorInternal.New(err.Error())
	}

	var subscriber models.Subscribers
	err = subscriberSnapshot.DataTo(&subscriber)
	if err != nil {
		return nil, ErrorInternal.New(err.Error())
	}

	return &subscriber, err
}

func PatchStructDataAndUpdate(ref *firestore.DocumentRef, oldData interface{}, newData interface{}) error {
	err := PatchStructData(oldData, newData)
	if err != nil {
		return err
	}

	_, err = ref.Set(context.Background(), oldData)

	if err != nil {
		return ErrorInternal.New(err.Error())
	}

	return nil
}

func PatchStructData(oldData interface{}, newData interface{}) error {
	newDataType := reflect.TypeOf(newData)
	newDataValue := reflect.ValueOf(newData)
	oldDataValue := reflect.ValueOf(oldData).Elem()

	if oldDataValue.Kind() != reflect.Struct || newDataValue.Kind() != reflect.Struct ||
		newDataType.Kind() != reflect.Struct {
		fmt.Println(oldDataValue.Kind(), newDataValue.Kind(), newDataType.Kind(), "data type")
		return ErrorInternal.New("patch data only accept struct")
	}

	for i := 0; i < newDataType.NumField(); i++ {
		newField := newDataType.Field(i)
		newFieldName := newField.Name
		newFieldValue := newDataValue.FieldByName(newFieldName)

		if newFieldValue.IsZero() {
			continue
		}

		oldFieldValue := oldDataValue.FieldByName(newFieldName)
		if !oldFieldValue.IsValid() {
			continue
		}

		if oldFieldValue.Kind() != newFieldValue.Kind() {
			fmt.Println("old field and new field value is not the same kind", oldFieldValue.Kind(),
				newFieldValue.Kind())
			continue
		}

		if !oldFieldValue.CanSet() {
			continue
		}
		oldFieldValue.Set(newFieldValue)
	}

	return nil
}
