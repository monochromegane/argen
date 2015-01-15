package ar

type limit struct {
	limit int
}

func (l *limit) setLimit(limit int) {
	l.limit = limit
}

func (l *limit) build() (query string, binds []interface{}) {
	return " LIMIT ?", []interface{}{l.limit}
}
