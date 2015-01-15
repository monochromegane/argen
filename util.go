package ar

import "reflect"

func IsZero(v interface{}) bool {
	return reflect.ValueOf(v).Interface() == reflect.Zero(reflect.TypeOf(v)).Interface()
}
