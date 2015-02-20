//go:generate go run ../cmd/argen/main.go
package tests

import "github.com/monochromegane/argen"

//+AR
type User struct {
	Id   int `db:"pk"`
	Name string
	Age  int
}

func (m User) scopeOlderThan(scope ar.Scope) *ar.Relation {
	return scope.Where("age", ">", scope.Args[0])
}
