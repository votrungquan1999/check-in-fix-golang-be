package utils

import (
	"cloud.google.com/go/firestore"
	"fmt"
	"google.golang.org/api/iterator"
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
