package query

type limit struct {
	limit int
}

func (l *limit) setLimit(limit int) {
	l.limit = limit
}

func (l *limit) build() (string, []interface{}) {
	return " LIMIT ?", []interface{}{l.limit}
}
