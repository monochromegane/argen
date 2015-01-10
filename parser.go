package goar

import (
	"go/ast"
	"go/token"
	"strings"
)

type Struct struct {
	Package    string
	Anotations []string
	Name       string
	Fields     []Field
}

type Field struct {
	Typ  string
	Name string
	Tag  string
}

func findAnotationComments(commments []*ast.Comment, anotation string) ([]string, bool) {
	result := []string{}
	hasAnotation := false
	for _, c := range commments {
		t := strings.TrimSpace(strings.TrimLeft(c.Text, "//"))
		if !strings.HasPrefix(t, anotation) {
			continue
		}
		hasAnotation = true
		result = append(result, t)
	}
	return result, hasAnotation
}

func AnotatedStructs(f *ast.File, anotation string) []Struct {

	result := []Struct{}

	pkg := f.Name.Name
	ast.Inspect(f, func(n ast.Node) bool {

		g, ok := n.(*ast.GenDecl)

		if !ok || g.Tok != token.TYPE {
			return true
		}

		st := Struct{Package: pkg}

		comments, hasAnotation := findAnotationComments(g.Doc.List, anotation)
		if !hasAnotation {
			return true
		}
		st.Anotations = comments

		for _, spec := range g.Specs {
			t := spec.(*ast.TypeSpec)
			st.Name = t.Name.Name

			s, ok := t.Type.(*ast.StructType)
			if !ok {
				return true
			}

			for _, f := range s.Fields.List {
				st.Fields = append(st.Fields, Field{
					Name: f.Names[0].Name,
					Typ:  f.Type.(*ast.Ident).Name,
					Tag:  f.Tag.Value,
				})
			}
		}
		result = append(result, st)
		return false
	})

	return result
}
