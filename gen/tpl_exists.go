package gen

var exists = &Template{
	Name: "Exists",
	Text: `
func (m {{.Name}}) Exists() bool {
	return m.newRelation().Exists()
}
`}
