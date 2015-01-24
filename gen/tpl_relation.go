package gen

var relation = &Template{
	Name: "Relation",
	Text: `
type {{.Name}}Relation struct {
	src *{{.Name}}
	*ar.Relation
}
`}
