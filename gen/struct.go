package gen

type structs []*structType

type structType struct {
	Package  string
	Comments comments
	Name     string
	Fields   []field
	Funcs    funcs
}

type comments []comment

type comment string

type field struct {
	Type string
	Name string
	Tag  tag
}

type tag string

type funcs []funcType

type funcType struct {
	Recv     string
	Comments comments
	Name     string
}
