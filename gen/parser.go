package gen

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

		comments, hasAnotation := findAnotatedComments(g.Doc.List, anotation)
		if !hasAnotation {
			return true
		}

		st, ok := findStruct(g.Specs)
		if !ok {
			return true
		}

		st.Anotations = comments
		st.pkg = pkg

		structs = append(structs, st)
		return false
	})

	return structs
}

func StructFuncs(f *ast.File) map[string]funcs {

	structFuncs := map[string]funcs{}

	ast.Inspect(f, func(n ast.Node) bool {

		f, ok := n.(*ast.FuncDecl)

		if !ok {
			return true
		}

		recv, ok := findRecv(f.Recv)
		if !ok {
			return true
		}

		fn := funcType{
			Recv:       recv,
			Name:       f.Name.Name,
			Anotations: findComments(f.Doc.List),
		}

		structFuncs[recv] = append(structFuncs[recv], fn)
		return false
	})

	return structFuncs
}

func findComments(commments []*ast.Comment) anotations {
	result := anotations{}
	for _, c := range commments {
		t := strings.TrimSpace(strings.TrimLeft(c.Text, "//"))
		result = append(result, anotation(t))
	}
	return result
}

func findAnotatedComments(commments []*ast.Comment, mark string) (anotations, bool) {
	result := anotations{}
	hasAnotation := false
	for _, c := range commments {
		t := strings.TrimSpace(strings.TrimLeft(c.Text, "//"))
		if strings.HasPrefix(t, mark) {
			hasAnotation = true
		}
		result = append(result, anotation(t))
	}
	return result, hasAnotation
}

func findStruct(specs []ast.Spec) (*structType, bool) {
	st := &structType{}
	for _, spec := range specs {
		t := spec.(*ast.TypeSpec)
		s, ok := t.Type.(*ast.StructType)
		if !ok {
			return st, false
		}

		st.Name = t.Name.Name
		for _, f := range s.Fields.List {
			field := field{
				Name: f.Names[0].Name,
				typ:  f.Type.(*ast.Ident).Name,
			}
			if f.Tag != nil {
				field.tag = tag(f.Tag.Value)
			}
			st.Fields = append(st.Fields, field)
		}
	}
	return st, true
}

func findRecv(recv *ast.FieldList) (string, bool) {
	if recv == nil {
		return "", false
	}
	for _, r := range recv.List {
		f, ok := r.Type.(*ast.Ident)
		if !ok {
			return "", false
		}
		return f.Name, true
	}
	return "", false
}
