package gen

var count = &Template{
	Name: "Count",
	Text: `
func (m {{.Name}}) Count(column ...string) int {
	return m.newRelation().Count(column...)
}
`}
