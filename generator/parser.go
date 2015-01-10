package generator

import (
	"go/ast"
	"go/token"
	"strings"
)

func AnotatedStructs(f *ast.File, anotation string) structs {

	structs := structs{}

	pkg := f.Name.Name
	ast.Inspect(f, func(n ast.Node) bool {

		g, ok := n.(*ast.GenDecl)

		if !ok || g.Tok != token.TYPE {
			return true
		}

		comments, hasAnotation := findComments(g.Doc.List, anotation)
		if !hasAnotation {
			return true
		}

		st, ok := findStruct(g.Specs)
		if !ok {
			return true
		}

		st.anotations = comments
		st.pkg = pkg

		structs = append(structs, st)
		return false
	})

	return structs
}

type structs []structType

func (s structs) pkg() string {
	return s[0].pkg
}

type structType struct {
	pkg        string
	anotations []string
	name       string
	fields     []field
}

type field struct {
	typ  string
	name string
	tag  string
}

func findComments(commments []*ast.Comment, anotation string) ([]string, bool) {
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

func findStruct(specs []ast.Spec) (structType, bool) {
	st := structType{}
	for _, spec := range specs {
		t := spec.(*ast.TypeSpec)
		s, ok := t.Type.(*ast.StructType)
		if !ok {
			return st, false
		}

		st.name = t.Name.Name
		for _, f := range s.Fields.List {
			st.fields = append(st.fields, field{
				name: f.Names[0].Name,
				typ:  f.Type.(*ast.Ident).Name,
				tag:  f.Tag.Value,
			})
		}
	}
	return st, true
}
