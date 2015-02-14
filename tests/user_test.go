package tests

import (
	"database/sql"
	"log"
	"os"
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
	expect := &User{Name: "test"}
	expect.Save()
	defer User{}.DeleteAll()

	u, err := User{}.Find(1)
	assertError(t, err)

	if u.Id != 1 {
		t.Errorf("column value should be equal to %v, but %v", 1, u.Id)
	}
	if expect.Name != u.Name {
		t.Errorf("column value should be equal to %v, but %v", expect.Name, u.Name)
	}
}

func assertError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("error should be nil, but %v", err)
	}
}
