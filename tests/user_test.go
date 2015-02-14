package tests

import (
	"database/sql"
	"log"
	"os"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/monochromegane/argen"
)

func TestMain(m *testing.M) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	Use(db)
	sqlStmt := "create table users (id integer not null primary key, name text);"
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err, sqlStmt)
	}

	os.Exit(m.Run())
}

func TestSelect(t *testing.T) {
	u := &User{Name: "test"}
	u.Save()
	defer User{}.DeleteAll()

	u, err := User{}.Select("id").First()
	assertError(t, err)

	if !ar.IsZero(u.Name) {
		t.Errorf("column value should be empty, but %s", u.Name)
	}
}

func TestFind(t *testing.T) {
	u := &User{Name: "test"}
	u.Save()
	defer User{}.DeleteAll()

	expect, _ := User{}.First()

	u, err := User{}.Find(1)
	assertError(t, err)
	assertEqualStruct(t, expect, u)
}

func TestFindBy(t *testing.T) {
	u := &User{Name: "test"}
	u.Save()
	defer User{}.DeleteAll()

	expect, _ := User{}.First()

	u, err := User{}.FindBy("name", "test")
	assertError(t, err)
	assertEqualStruct(t, expect, u)
}

func TestFirst(t *testing.T) {
	for _, name := range []string{"test1", "test2"} {
		u := &User{Name: name}
		u.Save()
	}
	defer User{}.DeleteAll()

	expect, _ := User{}.Where("name", "test1").QueryRow()

	u, err := User{}.First()
	assertError(t, err)
	assertEqualStruct(t, expect, u)
}

func TestLast(t *testing.T) {
	for _, name := range []string{"test1", "test2"} {
		u := &User{Name: name}
		u.Save()
	}
	defer User{}.DeleteAll()

	expect, _ := User{}.Where("name", "test2").QueryRow()

	u, err := User{}.Last()
	assertError(t, err)
	assertEqualStruct(t, expect, u)
}

func assertEqualStruct(t *testing.T, expect, actual interface{}) {
	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("struct should be equal to %v, but %v", expect, actual)
	}
}

func assertError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("error should be nil, but %v", err)
	}
}
