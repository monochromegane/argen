package tests

import (
	"fmt"

	"github.com/monochromegane/argen"
	"github.com/monochromegane/goban"
)

type UserRelation struct {
	src *User
	*ar.Relation
}

func (m *User) newRelation() *UserRelation {
	r := ar.NewRelation()
	r.Table("users").Columns(
		"id",
		"name",
	)

	return &UserRelation{m, r}
}

func (m User) Select(columns ...string) *UserRelation {
	return m.newRelation().Select(columns...)
}

func (r *UserRelation) Select(columns ...string) *UserRelation {
	r.Relation.Columns(columns...)
	return r
}

func (m User) Find(id int) (*User, error) {
	return m.newRelation().Find(id)
}

func (r *UserRelation) Find(id int) (*User, error) {
	return r.FindBy("id", id)
}

func (m User) FindBy(cond string, args ...interface{}) (*User, error) {
	return m.newRelation().FindBy(cond, args...)
}

func (r *UserRelation) FindBy(cond string, args ...interface{}) (*User, error) {
	return r.Where(cond, args...).Limit(1).QueryRow()
}

func (m User) First() (*User, error) {
	return m.newRelation().First()
}

func (r *UserRelation) First() (*User, error) {
	return r.Order("id", "ASC").Limit(1).QueryRow()
}

func (m User) Last() (*User, error) {
	return m.newRelation().Last()
}

func (r *UserRelation) Last() (*User, error) {
	return r.Order("id", "DESC").Limit(1).QueryRow()
}

func (m User) Where(cond string, args ...interface{}) *UserRelation {
	return m.newRelation().Where(cond, args...)
}

func (r *UserRelation) Where(cond string, args ...interface{}) *UserRelation {
	r.Relation.Where(cond, args...)
	return r
}

func (r *UserRelation) And(cond string, args ...interface{}) *UserRelation {
	r.Relation.And(cond, args...)
	return r
}

func (m User) Order(column, order string) *UserRelation {
	return m.newRelation().Order(column, order)
}

func (r *UserRelation) Order(column, order string) *UserRelation {
	r.Relation.OrderBy(column, order)
	return r
}

func (m User) Limit(limit int) *UserRelation {
	return m.newRelation().Limit(limit)
}

func (r *UserRelation) Limit(limit int) *UserRelation {
	r.Relation.Limit(limit)
	return r
}

func (m User) Offset(offset int) *UserRelation {
	return m.newRelation().Offset(offset)
}

func (r *UserRelation) Offset(offset int) *UserRelation {
	r.Relation.Offset(offset)
	return r
}

func (m User) Group(group string, groups ...string) *UserRelation {
	return m.newRelation().Group(group, groups...)
}

func (r *UserRelation) Group(group string, groups ...string) *UserRelation {
	r.Relation.GroupBy(group, groups...)
	return r
}

func (r *UserRelation) Having(cond string, args ...interface{}) *UserRelation {
	r.Relation.Having(cond, args...)
	return r
}

func (r *UserRelation) Explain() error {
	r.Relation.Explain()
	q, b := r.Build()
	rows, err := db.Query(q, b...)
	if err != nil {
		return err
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	var values [][]string
	for rows.Next() {
		vals := make([]string, len(columns))
		ptrs := make([]interface{}, len(columns))
		for i, _ := range vals {
			ptrs[i] = &vals[i]
		}
		rows.Scan(ptrs...)
		values = append(values, vals)
	}

	fmt.Printf("%s %v\n", q, b)
	goban.Render(columns, values)
	return nil
}

func (m User) IsValid() (bool, *ar.Errors) {
	result := true
	errors := &ar.Errors{}
	rules := map[string]*ar.Validation{}
	for name, rule := range rules {
		if ok, errs := ar.NewValidator(rule).IsValid(m.fieldValueByName(name)); !ok {
			result = false
			errors.SetErrors(name, errs)
		}
	}
	customs := []ar.CustomValidator{}
	for _, c := range customs {
		c(errors)
	}
	return result, errors
}

type UserParams User

func (m User) Create(p UserParams) (*User, *ar.Errors) {
	n := &User{
		Id:   p.Id,
		Name: p.Name,
	}
	_, errs := n.Save()
	return n, errs
}

func (m *User) IsNewRecord() bool {
	return ar.IsZero(m.Id)
}

func (m *User) IsPersistent() bool {
	return !m.IsNewRecord()
}

func (m *User) Save() (bool, *ar.Errors) {
	if ok, errs := m.IsValid(); !ok {
		return false, errs
	}
	errs := &ar.Errors{}
	if m.IsNewRecord() {
		ins := ar.NewInsert()
		q, b := ins.Table("users").Params(map[string]interface{}{
			"name": m.Name,
		}).Build()

		if result, err := db.Exec(q, b...); err != nil {
			errs.AddError("base", err)
			return false, errs
		} else {
			if lastId, err := result.LastInsertId(); err == nil {
				m.Id = int(lastId)
			}
		}
		return true, nil
	} else {
		params := map[string]interface{}{
			"id":   m.Id,
			"name": m.Name,
		}
		if _, err := m.updateColumnsByMap(params); err != nil {
			errs.AddError("base", err)
			return false, errs
		}
		return true, nil
	}
}

func (m *User) Update(p UserParams) (bool, *ar.Errors) {
	if ok, errs := m.IsValid(); !ok {
		return false, errs
	}
	return m.UpdateColumns(p)
}

func (m *User) UpdateColumns(p UserParams) (bool, *ar.Errors) {
	errs := &ar.Errors{}
	params := map[string]interface{}{}

	if !ar.IsZero(p.Id) && m.Id != p.Id {
		params["id"] = p.Id
	}

	if !ar.IsZero(p.Name) && m.Name != p.Name {
		params["name"] = p.Name
	}

	if _, err := m.updateColumnsByMap(params); err != nil {
		errs.AddError("base", err)
		return false, errs
	}
	return true, nil
}

func (m *User) updateColumnsByMap(params map[string]interface{}) (bool, error) {
	upd := ar.NewUpdate()
	q, b := upd.Table("users").Params(params).Where("id", m.Id).Build()
	if _, err := db.Exec(q, b...); err != nil {
		return false, err
	}
	return true, nil
}

func (m User) DeleteAll() (bool, *ar.Errors) {
	errs := &ar.Errors{}
	del := ar.NewDelete()
	del.Table("users")
	q, b := del.Build()
	if _, err := db.Exec(q, b...); err != nil {
		errs.AddError("base", err)
		return false, errs
	}
	return true, nil
}

func (r *UserRelation) Query() ([]*User, error) {
	q, b := r.Build()
	rows, err := db.Query(q, b...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []*User{}
	for rows.Next() {
		row := &User{}
		err := rows.Scan(row.fieldPtrsByName(r.Relation.GetColumns())...)
		if err != nil {
			return nil, err
		}
		results = append(results, row)
	}
	return results, nil
}

func (r *UserRelation) QueryRow() (*User, error) {
	q, b := r.Build()
	row := &User{}
	err := db.QueryRow(q, b...).Scan(row.fieldPtrsByName(r.Relation.GetColumns())...)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (m *User) fieldValueByName(name string) interface{} {
	switch name {
	case "id":
		return m.Id
	case "name":
		return m.Name
	default:
		return ""
	}
}

func (m *User) fieldPtrByName(name string) interface{} {
	switch name {
	case "id":
		return &m.Id
	case "name":
		return &m.Name
	default:
		return nil
	}
}

func (m *User) fieldPtrsByName(names []string) []interface{} {
	fields := []interface{}{}
	for _, n := range names {
		f := m.fieldPtrByName(n)
		fields = append(fields, f)
	}
	return fields
}
