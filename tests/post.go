//go:generate go run ../cmd/argen/main.go
package tests

import "github.com/monochromegane/argen"

//+AR
type Post struct {
	Id     int `db:"pk"`
	UserId int `db:"fk"`
	Name   string
}

func (p Post) validatesName() ar.Rule {
	return ar.MakeRule().Format().With("/name/")
}

func (p Post) validateCustom() ar.Rule {
	return ar.CustomRule(func(errors *ar.Errors) {
		if p.Name != "name" {
			errors.Add("name", "must be name")
		}
	})
}
