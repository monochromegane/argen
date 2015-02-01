package gen

var belongsTo = &Template{
	Name: "BelongsTo",
	Text: `
func (m *{{.Recv.Name}}) {{.Func}}() (*{{.Model}}, error) {
	asc := m.{{.FuncName}}()
	pk := "{{.PrimaryKey}}"
	fk := "{{.ForeignKey}}"
	if asc != nil && asc.PrimaryKey != "" {
		pk = asc.PrimaryKey
	}
	if asc != nil && asc.ForeignKey != "" {
		fk = asc.ForeignKey
	}
	return {{.Model}}{}.Where(pk, m.fieldByName(ar.ToCamelCase(fk))).QueryRow()
}
`}
