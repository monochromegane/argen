package gen

type field struct {
	Type string
	Name string
	Tag  tag
}

func (f field) ColumnName() string {
	return toSnakeCase(f.Name)
}

func (f field) isPrimaryKey() bool {
	if pk := f.Tag.get("db"); pk == "pk" {
		return true
	}
	return false
}
