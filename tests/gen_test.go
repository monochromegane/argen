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
	sqlStmt := `
	create table users (id integer not null primary key, name text);
	create table posts (id integer not null primary key, user_id integer, name text);
	`
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
	assertEqualStruct(t, expect, u)
}

func TestFindBy(t *testing.T) {
	expect := &User{Name: "test"}
	expect.Save()
	defer User{}.DeleteAll()

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

func TestWhere(t *testing.T) {
	expect := &User{Name: "test"}
	expect.Save()
	defer User{}.DeleteAll()

	u, err := User{}.Where("name", "test").And("id", 1).QueryRow()

	assertError(t, err)
	assertEqualStruct(t, expect, u)
}

func TestOrder(t *testing.T) {
	expects := []string{"test1", "test2"}
	for _, name := range expects {
		u := &User{Name: name}
		u.Save()
	}
	defer User{}.DeleteAll()

	users, err := User{}.Order("name", "ASC").Query()

	assertError(t, err)
	for i, u := range users {
		if u.Name != expects[i] {
			t.Errorf("column value should be %v, but %v", expects[i], u.Name)
		}
	}
}

func TestLimitAndOffset(t *testing.T) {
	for _, name := range []string{"test1", "test2", "test3"} {
		u := &User{Name: name}
		u.Save()
	}
	defer User{}.DeleteAll()

	users, err := User{}.Limit(2).Offset(1).Order("name", "ASC").Query()

	assertError(t, err)
	expects := []string{"test2", "test3"}
	for i, u := range users {
		if u.Name != expects[i] {
			t.Errorf("column value should be %v, but %v", expects[i], u.Name)
		}
	}
}

func TestGroupByAndHaving(t *testing.T) {
	for _, name := range []string{"testA", "testB", "testB"} {
		u := &User{Name: name}
		u.Save()
	}
	defer User{}.DeleteAll()

	users, err := User{}.Group("name").Having("count(name)", 2).Query()

	assertError(t, err)
	expects := []string{"testB"}
	for i, u := range users {
		if u.Name != expects[i] {
			t.Errorf("column value should be %v, but %v", expects[i], u.Name)
		}
	}
}

func TestExplain(t *testing.T) {
	err := User{}.Where("name", "test").Explain()
	assertError(t, err)
}

func TestIsValid(t *testing.T) {
	p := &Post{Name: "abc"}
	_, errs := p.IsValid()

	if len(errs.Messages["name"]) != 1 {
		t.Errorf("errors count should be 1, but %d", len(errs.Messages["name"]))
	}
}

func TestCreate(t *testing.T) {
	u, errs := User{}.Create(UserParams{
		Name: "TestCreate",
	})
	defer User{}.DeleteAll()

	assertErrors(t, errs)

	expect, _ := User{}.FindBy("name", "TestCreate")
	assertEqualStruct(t, expect, u)
}

func TestIsNewRecordAndIsPresistent(t *testing.T) {
	defer User{}.DeleteAll()

	u := &User{Name: "test"}
	if !u.IsNewRecord() {
		t.Errorf("struct is new record, but isn't new record.")
	}

	u.Save()
	if !u.IsPersistent() {
		t.Errorf("struct is persistent, but isn't persistent.")
	}
}

func TestSaveWithInvalidData(t *testing.T) {
	defer Post{}.DeleteAll()

	// OnCreate
	p := &Post{Name: "invalid"}
	_, errs := p.Save()

	if len(errs.Messages["name"]) != 1 {
		t.Errorf("errors count should be 1, but %d", len(errs.Messages["name"]))
	}

	p.Name = "name"
	_, errs = p.Save()
	assertErrors(t, errs)

	// OnUpdate
	p.Name = "invalid2"
	_, errs = p.Save()

	if len(errs.Messages["name"]) != 1 {
		t.Errorf("errors count should be 1, but %d", len(errs.Messages["name"]))
	}

}

func TestSave(t *testing.T) {
	defer User{}.DeleteAll()

	u := &User{Name: "test"}

	_, errs := u.Save()
	assertErrors(t, errs)

	if u.Id == 0 {
		t.Errorf("Id should be setted after save, but isn't setted")
	}

	expect, _ := User{}.FindBy("name", "test")
	assertEqualStruct(t, expect, u)

	u.Name = "test2"
	_, errs = u.Save()
	assertErrors(t, errs)

	expect, _ = User{}.Find(u.Id)
	assertEqualStruct(t, expect, u)
}

func TestUpdate(t *testing.T) {
	defer User{}.DeleteAll()

	u := &User{Name: "test"}
	_, errs := u.Save()

	expect := UserParams{Name: "test2"}
	_, errs = u.Update(expect)
	assertErrors(t, errs)

	actual, _ := User{}.Find(u.Id)
	if expect.Name != actual.Name {
		t.Errorf("column value should be equal to %v, but %v", expect.Name, actual.Name)
	}
}

func TestUpdateColumns(t *testing.T) {
	defer User{}.DeleteAll()

	u := &User{Name: "test"}
	_, errs := u.Save()

	expect := UserParams{Name: "test2"}
	_, errs = u.UpdateColumns(expect)
	assertErrors(t, errs)

	actual, _ := User{}.Find(u.Id)
	if expect.Name != actual.Name {
		t.Errorf("column value should be equal to %v, but %v", expect.Name, actual.Name)
	}
}

func assertEqualStruct(t *testing.T, expect, actual interface{}) {
	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("struct should be equal to %v, but %v", expect, actual)
	}
}

func assertErrors(t *testing.T, errs *ar.Errors) {
	if errs != nil {
		t.Errorf("errors should be nil, but %v", errs)
	}
}

func assertError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("error should be nil, but %v", err)
	}
}
