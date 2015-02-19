package tests

import (
	"fmt"

	"github.com/monochromegane/argen"
	"github.com/monochromegane/goban"
)

type PostRelation struct {
	src *Post
	*ar.Relation
}

func (m *Post) newRelation() *PostRelation {
	r := ar.NewRelation()
	r.Table("posts").Columns(
		"id",
		"user_id",
		"name",
	)

	return &PostRelation{m, r}
}

func (m Post) Select(columns ...string) *PostRelation {
	return m.newRelation().Select(columns...)
}

func (r *PostRelation) Select(columns ...string) *PostRelation {
	r.Relation.Columns(columns...)
	return r
}

func (m Post) Find(id int) (*Post, error) {
	return m.newRelation().Find(id)
}

func (r *PostRelation) Find(id int) (*Post, error) {
	return r.FindBy("id", id)
}

func (m Post) FindBy(cond string, args ...interface{}) (*Post, error) {
	return m.newRelation().FindBy(cond, args...)
}

func (r *PostRelation) FindBy(cond string, args ...interface{}) (*Post, error) {
	return r.Where(cond, args...).Limit(1).QueryRow()
}

func (m Post) First() (*Post, error) {
	return m.newRelation().First()
}

func (r *PostRelation) First() (*Post, error) {
	return r.Order("id", "ASC").Limit(1).QueryRow()
}

func (m Post) Last() (*Post, error) {
	return m.newRelation().Last()
}

func (r *PostRelation) Last() (*Post, error) {
	return r.Order("id", "DESC").Limit(1).QueryRow()
}

func (m Post) Where(cond string, args ...interface{}) *PostRelation {
	return m.newRelation().Where(cond, args...)
}

func (r *PostRelation) Where(cond string, args ...interface{}) *PostRelation {
	r.Relation.Where(cond, args...)
	return r
}

func (r *PostRelation) And(cond string, args ...interface{}) *PostRelation {
	r.Relation.And(cond, args...)
	return r
}

func (m Post) Order(column, order string) *PostRelation {
	return m.newRelation().Order(column, order)
}

func (r *PostRelation) Order(column, order string) *PostRelation {
	r.Relation.OrderBy(column, order)
	return r
}

func (m Post) Limit(limit int) *PostRelation {
	return m.newRelation().Limit(limit)
}

func (r *PostRelation) Limit(limit int) *PostRelation {
	r.Relation.Limit(limit)
	return r
}

func (m Post) Offset(offset int) *PostRelation {
	return m.newRelation().Offset(offset)
}

func (r *PostRelation) Offset(offset int) *PostRelation {
	r.Relation.Offset(offset)
	return r
}

func (m Post) Group(group string, groups ...string) *PostRelation {
	return m.newRelation().Group(group, groups...)
}

func (r *PostRelation) Group(group string, groups ...string) *PostRelation {
	r.Relation.GroupBy(group, groups...)
	return r
}

func (r *PostRelation) Having(cond string, args ...interface{}) *PostRelation {
	r.Relation.Having(cond, args...)
	return r
}

func (r *PostRelation) Explain() error {
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

func (m Post) IsValid() (bool, *ar.Errors) {
	result := true
	errors := &ar.Errors{}
	var on ar.On
	if m.IsNewRecord() {
		on = ar.OnCreate()
	} else {
		on = ar.OnUpdate()
	}
	rules := map[string]*ar.Validation{
		"name": m.validatesName().Rule(),
	}
	for name, rule := range rules {
		if ok, errs := ar.NewValidator(rule).On(on).IsValid(m.fieldValueByName(name)); !ok {
			result = false
			errors.SetErrors(name, errs)
		}
	}
	customs := []*ar.Validation{
		m.validateCustom().Rule(),
	}
	for _, rule := range customs {
		custom := ar.NewValidator(rule).On(on).Custom()
		custom(errors)
	}
	if len(errors.Messages) > 0 {
		result = false
	}
	return result, errors
}

type PostParams Post

func (m Post) Create(p PostParams) (*Post, *ar.Errors) {
	n := &Post{
		Id:     p.Id,
		UserId: p.UserId,
		Name:   p.Name,
	}
	_, errs := n.Save()
	return n, errs
}

func (m *Post) IsNewRecord() bool {
	return ar.IsZero(m.Id)
}

func (m *Post) IsPersistent() bool {
	return !m.IsNewRecord()
}

func (m *Post) Save(validate ...bool) (bool, *ar.Errors) {
	if len(validate) == 0 || len(validate) > 0 && validate[0] {
		if ok, errs := m.IsValid(); !ok {
			return false, errs
		}
	}
	errs := &ar.Errors{}
	if m.IsNewRecord() {
		ins := ar.NewInsert()
		q, b := ins.Table("posts").Params(map[string]interface{}{
			"user_id": m.UserId,
			"name":    m.Name,
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
		upd := ar.NewUpdate()
		q, b := upd.Table("posts").Params(map[string]interface{}{
			"id":      m.Id,
			"user_id": m.UserId,
			"name":    m.Name,
		}).Where("id", m.Id).Build()
		if _, err := db.Exec(q, b...); err != nil {
			errs.AddError("base", err)
			return false, errs
		}
		return true, nil
	}
}

func (m *Post) Update(p PostParams) (bool, *ar.Errors) {

	if !ar.IsZero(p.Id) {
		m.Id = p.Id
	}
	if !ar.IsZero(p.UserId) {
		m.UserId = p.UserId
	}
	if !ar.IsZero(p.Name) {
		m.Name = p.Name
	}
	return m.Save()
}

func (m *Post) UpdateColumns(p PostParams) (bool, *ar.Errors) {

	if !ar.IsZero(p.Id) {
		m.Id = p.Id
	}
	if !ar.IsZero(p.UserId) {
		m.UserId = p.UserId
	}
	if !ar.IsZero(p.Name) {
		m.Name = p.Name
	}
	return m.Save(false)
}

func (m *Post) Destroy() (bool, *ar.Errors) {
	return m.Delete()
}

func (m *Post) Delete() (bool, *ar.Errors) {
	errs := &ar.Errors{}
	del := ar.NewDelete()
	q, b := del.Table("posts").Where("id", m.Id).Build()
	if _, err := db.Exec(q, b...); err != nil {
		errs.AddError("base", err)
		return false, errs
	}
	return true, nil
}

func (m Post) DeleteAll() (bool, *ar.Errors) {
	errs := &ar.Errors{}
	del := ar.NewDelete()
	del.Table("posts")
	q, b := del.Build()
	if _, err := db.Exec(q, b...); err != nil {
		errs.AddError("base", err)
		return false, errs
	}
	return true, nil
}

func (r *PostRelation) Query() ([]*Post, error) {
	q, b := r.Build()
	rows, err := db.Query(q, b...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []*Post{}
	for rows.Next() {
		row := &Post{}
		err := rows.Scan(row.fieldPtrsByName(r.Relation.GetColumns())...)
		if err != nil {
			return nil, err
		}
		results = append(results, row)
	}
	return results, nil
}

func (r *PostRelation) QueryRow() (*Post, error) {
	q, b := r.Build()
	row := &Post{}
	err := db.QueryRow(q, b...).Scan(row.fieldPtrsByName(r.Relation.GetColumns())...)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (m *Post) fieldValueByName(name string) interface{} {
	switch name {
	case "id":
		return m.Id
	case "user_id":
		return m.UserId
	case "name":
		return m.Name
	default:
		return ""
	}
}

func (m *Post) fieldPtrByName(name string) interface{} {
	switch name {
	case "id":
		return &m.Id
	case "user_id":
		return &m.UserId
	case "name":
		return &m.Name
	default:
		return nil
	}
}

func (m *Post) fieldPtrsByName(names []string) []interface{} {
	fields := []interface{}{}
	for _, n := range names {
		f := m.fieldPtrByName(n)
		fields = append(fields, f)
	}
	return fields
}
