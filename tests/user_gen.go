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
	r := ar.NewRelation(db, logger)
	r.Table("users").Columns(
		"id",
		"name",
		"age",
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
	rows, err := r.Relation.Explain().Query()
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

	goban.Render(columns, values)
	return nil
}

func (m User) IsValid() (bool, *ar.Errors) {
	result := true
	errors := &ar.Errors{}
	var on ar.On
	if m.IsNewRecord() {
		on = ar.OnCreate()
	} else {
		on = ar.OnUpdate()
	}
	rules := map[string]*ar.Validation{}
	for name, rule := range rules {
		if ok, errs := ar.NewValidator(rule).On(on).IsValid(m.fieldValueByName(name)); !ok {
			result = false
			errors.SetErrors(name, errs)
		}
	}
	customs := []*ar.Validation{}
	for _, rule := range customs {
		custom := ar.NewValidator(rule).On(on).Custom()
		custom(errors)
	}
	if len(errors.Messages) > 0 {
		result = false
	}
	return result, errors
}

func (m User) OlderThan(args ...interface{}) *UserRelation {
	r := m.newRelation()
	m.scopeOlderThan(ar.Scope{r.Relation, args})
	return r
}

func (r *UserRelation) OlderThan(args ...interface{}) *UserRelation {
	r.src.scopeOlderThan(ar.Scope{r.Relation, args})
	return r
}

func (m *User) Posts() ([]*Post, error) {
	asc := m.hasManyPosts()
	fk := "user_id"
	if asc != nil && asc.ForeignKey != "" {
		fk = asc.ForeignKey
	}
	return Post{}.Where(fk, m.Id).Query()
}

func (m User) JoinsPosts() *UserRelation {
	return m.newRelation().JoinsPosts()
}

func (r *UserRelation) JoinsPosts() *UserRelation {
	asc := r.src.hasManyPosts()
	fk := "user_id"
	if asc != nil && asc.ForeignKey != "" {
		fk = asc.ForeignKey
	}
	r.Relation.InnerJoin("posts", fmt.Sprintf("posts.%s = users.id", fk))
	return r
}

func (m *User) BuildPost(p PostParams) *Post {
	p.UserId = m.Id
	return Post{}.Build(p)
}

type UserParams User

func (m User) Build(p UserParams) *User {
	return &User{
		Id:   p.Id,
		Name: p.Name,
		Age:  p.Age,
	}
}

func (m User) Create(p UserParams) (*User, *ar.Errors) {
	n := m.Build(p)
	_, errs := n.Save()
	return n, errs
}

func (m *User) IsNewRecord() bool {
	return ar.IsZero(m.Id)
}

func (m *User) IsPersistent() bool {
	return !m.IsNewRecord()
}

func (m *User) Save(validate ...bool) (bool, *ar.Errors) {
	if len(validate) == 0 || len(validate) > 0 && validate[0] {
		if ok, errs := m.IsValid(); !ok {
			return false, errs
		}
	}
	errs := &ar.Errors{}
	if m.IsNewRecord() {
		ins := ar.NewInsert(db, logger).Table("users").Params(map[string]interface{}{
			"name": m.Name,
			"age":  m.Age,
		})

		if result, err := ins.Exec(); err != nil {
			errs.AddError("base", err)
			return false, errs
		} else {
			if lastId, err := result.LastInsertId(); err == nil {
				m.Id = int(lastId)
			}
		}
		return true, nil
	} else {
		upd := ar.NewUpdate(db, logger).Table("users").Params(map[string]interface{}{
			"id":   m.Id,
			"name": m.Name,
			"age":  m.Age,
		}).Where("id", m.Id)

		if _, err := upd.Exec(); err != nil {
			errs.AddError("base", err)
			return false, errs
		}
		return true, nil
	}
}

func (m *User) Update(p UserParams) (bool, *ar.Errors) {

	if !ar.IsZero(p.Id) {
		m.Id = p.Id
	}
	if !ar.IsZero(p.Name) {
		m.Name = p.Name
	}
	if !ar.IsZero(p.Age) {
		m.Age = p.Age
	}
	return m.Save()
}

func (m *User) UpdateColumns(p UserParams) (bool, *ar.Errors) {

	if !ar.IsZero(p.Id) {
		m.Id = p.Id
	}
	if !ar.IsZero(p.Name) {
		m.Name = p.Name
	}
	if !ar.IsZero(p.Age) {
		m.Age = p.Age
	}
	return m.Save(false)
}

func (m *User) Destroy() (bool, *ar.Errors) {
	return m.Delete()
}

func (m *User) Delete() (bool, *ar.Errors) {
	errs := &ar.Errors{}
	if _, err := ar.NewDelete(db, logger).Table("users").Where("id", m.Id).Exec(); err != nil {
		errs.AddError("base", err)
		return false, errs
	}
	return true, nil
}

func (m User) DeleteAll() (bool, *ar.Errors) {
	errs := &ar.Errors{}
	if _, err := ar.NewDelete(db, logger).Table("users").Exec(); err != nil {
		errs.AddError("base", err)
		return false, errs
	}
	return true, nil
}

func (r *UserRelation) Query() ([]*User, error) {
	rows, err := r.Relation.Query()
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
	row := &User{}
	err := r.Relation.QueryRow(row.fieldPtrsByName(r.Relation.GetColumns())...)
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
	case "age":
		return m.Age
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
	case "age":
		return &m.Age
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
