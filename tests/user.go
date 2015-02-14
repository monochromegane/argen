//go:generate go run ../cmd/argen/main.go
package tests

//+AR
type User struct {
	Id   int `db:"pk"`
	Name string
}
