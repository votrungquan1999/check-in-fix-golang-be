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
