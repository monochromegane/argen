package gen

var destroy = &Template{
	Name: "Destroy",
	Text: `
func (m *{{.Name}}) Destroy() (bool, error) {
	return m.Delete()
}
`}
