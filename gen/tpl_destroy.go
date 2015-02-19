package gen

var destroy = &Template{
	Name: "Destroy",
	Text: `
func (m *{{.Name}}) Destroy() (bool, *ar.Errors) {
	return m.Delete()
}
`}
