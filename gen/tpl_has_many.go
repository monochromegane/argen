package gen

var hasMany = &Template{
	Name: "HasMany",
	Text: `
func (m *{{.Recv.Name}}) {{.Func}}() ([]*{{.Model}}, error) {
	asc := m.{{.FuncName}}()
	fk := "{{.ForeignKey}}"
	if asc != nil && asc.ForeignKey != "" {
		fk = asc.ForeignKey
	}
	return {{.Model}}{}.Where(fk, m.{{.Recv.PrimaryKeyField}}).Query()
}
`}
