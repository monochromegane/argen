package ar

import (
	"reflect"
	"strings"
	"unicode"

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

func ToCamelCase(s string) string {
	var camel string
	for _, split := range strings.Split(s, "_") {
		c := []rune(split)
		c[0] = unicode.ToUpper(c[0])
		camel += string(c)
	}
	return camel
}
