package utils

import (
	"reflect"
)

func InterfaceSlice(slice interface{}) []interface{} {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() != reflect.Slice && sliceVal.Kind() != reflect.Array {
		panic("InterfaceSlice expect to receive a slice or an array")
	}

	newSlice := make([]interface{}, 0)
	for index := 0; index < sliceVal.Len(); index++ {
		newSlice = append(newSlice, sliceVal.Index(index).Interface())
	}

	return newSlice
}

func SliceContain(slice []interface{}, value interface{}) bool {
	for _, arrValue := range slice {
		if reflect.DeepEqual(arrValue, value) {
			return true
		}
	}

	return false
}
