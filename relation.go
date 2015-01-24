package ar

import (
	"reflect"

	"github.com/monochromegane/argen/query"
)

type Relation struct {
	*query.Select
}

type Insert struct {
	*query.Insert
}

type Update struct {
	*query.Update
}

func IsZero(v interface{}) bool {
	return reflect.ValueOf(v).Interface() == reflect.Zero(reflect.TypeOf(v)).Interface()
}
