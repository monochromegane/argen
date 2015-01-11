package generator

import (
	"go/ast"
	"go/token"
	"regexp"
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

func (s structs) Package() string {
	if len(s) == 0 {
		return ""
	}
	return s[0].pkg
}

type structType struct {
	pkg        string
	anotations []string
	Name       string
	Fields     []field
}

func (s structType) FieldNames(prefix string) []string {
	names := []string{}
	for _, f := range s.Fields {
		names = append(names, prefix+f.Name)
	}
	return names
}

func (s structType) ColumnNames() []string {
	names := []string{}
	for _, f := range s.Fields {
		names = append(names, f.ColumnName())
	}
	return names
}

func (s structType) PrimaryKey() string {
	k, _ := s.primaryKey()
	return k
}

func (s structType) PrimaryKeyType() string {
	_, t := s.primaryKey()
	return t
}

func (s structType) primaryKey() (string, string) {
	for _, f := range s.Fields {
		if f.isPrimaryKey() {
			return f.Name, f.typ
		}
	}
	return "id", "int"
}

type field struct {
	typ  string
	Name string
	tag  tag
}

func (f field) ColumnName() string {
	if tag := f.tag.get("column"); tag != "" {
		return tag
	}
	return f.toSnakeCase()
}

func (f field) isPrimaryKey() bool {
	if pk := f.tag.get("db"); pk == "pk" {
		return true
	}
	return false
}

type tag string

func (t tag) tags() []string {
	return strings.Split(string(t), " ")
}

func (t tag) get(key string) string {
	var value string
	for _, tag := range t.tags() {
		if !strings.HasPrefix(tag, key+":") {
			continue
		}
		pair := strings.SplitN(tag, ":", 1)
		if !strings.HasPrefix(pair[1], "\"") || !strings.HasSuffix(pair[1], "\"") {
			continue
		}
		return strings.Trim(pair[1], "\"")
	}
	return value
}

func (f field) toSnakeCase() string {
	const snake = "${1}_${2}"
	reg1 := regexp.MustCompile("([A-Z]+)([A-Z][a-z])")
	reg2 := regexp.MustCompile("([a-z])([A-Z])")
	return strings.ToLower(reg2.ReplaceAllString(reg1.ReplaceAllString(f.Name, snake), snake))
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
