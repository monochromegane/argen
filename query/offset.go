package query

type offset struct {
	offset int
}

func (o *offset) setOffset(offset int) {
	o.offset = offset
}

func (o *offset) build() (string, []interface{}) {
	return " OFFSET ?", []interface{}{o.offset}
}
