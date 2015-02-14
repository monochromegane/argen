package gen

var delete = &Template{
	Name: "Delete",
	Text: `
func (m {{.Name}}) DeleteAll() (bool, *ar.Errors) {
        errs := &ar.Errors{}
        del := ar.NewDelete()
        del.Table("{{.TableName}}")
        q, b := del.Build()
        if _, err := db.Exec(q, b...); err != nil {
                errs.Add("base", err)
                return false, errs
        }
        return true, nil
}
`}
