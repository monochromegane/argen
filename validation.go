package ar

type Rule interface {
	Rule() *Validation
}

func MakeRule() *Validation {
	return &Validation{}
}

type Validation struct {
	presence     bool
	length       *length
	format       *format
	numericality *numericality
	inclusion    []string
	exclusion    []string
}

func (v *Validation) Rule() *Validation {
	return v
}

func (v *Validation) Presence() *Validation {
	v.presence = true
	return v
}

func (v *Validation) Length() *length {
	if v.length == nil {
		v.length = &length{Validation: v}
	}
	return v.length
}

func (v *Validation) Format() *format {
	if v.format == nil {
		v.format = &format{Validation: v}
	}
	return v.format
}

func (v *Validation) Numericality() *numericality {
	if v.numericality == nil {
		v.numericality = &numericality{Validation: v}
	}
	return v.numericality
}

func (v *Validation) Inclusion(collection ...string) *Validation {
	inclusion := []string{}
	v.inclusion = append(inclusion, collection...)
	return v
}

func (v *Validation) Exclusion(collection ...string) *Validation {
	exclusion := []string{}
	v.exclusion = append(exclusion, collection...)
	return v
}

type length struct {
	*Validation
	minimum  int
	maximum  int
	is       int
	from, to int
}

func (l *length) Rule() *Validation {
	return l.Validation
}

func (l *length) Minimum(minimum int) *length {
	l.minimum = minimum
	return l
}

func (l *length) Maximum(maximum int) *length {
	l.maximum = maximum
	return l
}

func (l *length) Is(is int) *length {
	l.is = is
	return l
}

func (l *length) In(from, to int) *length {
	l.from = from
	l.to = to
	return l
}

func (l *length) WithIn(from, to int) *length {
	return l.In(from, to)
}

type format struct {
	*Validation
	with string
}

func (f *format) Rule() *Validation {
	return f.Validation
}

func (f *format) With(regexp string) *format {
	f.with = regexp
	return f
}

type numericality struct {
	*Validation
	onlyInteger          bool
	greaterThan          int
	greaterThanOrEqualTo int
	equalTo              int
	lessThan             int
	lessThanOrEqualTo    int
	odd                  bool
	even                 bool
}

func (n *numericality) OnlyInteger() *numericality {
	n.onlyInteger = true
	return n
}

func (n *numericality) GreaterThan(num int) *numericality {
	n.greaterThan = num
	return n
}

func (n *numericality) GreaterThanOrEqualTo(num int) *numericality {
	n.greaterThanOrEqualTo = num
	return n
}

func (n *numericality) EqualTo(num int) *numericality {
	n.equalTo = num
	return n
}

func (n *numericality) LessThan(num int) *numericality {
	n.lessThan = num
	return n
}

func (n *numericality) LessThanOrEqualTo(num int) *numericality {
	n.lessThanOrEqualTo = num
	return n
}

func (n *numericality) Odd() *numericality {
	n.odd = true
	return n
}

func (n *numericality) Even() *numericality {
	n.even = true
	return n
}

type inclusion struct {
	*Validation
	in []string
}

func (i *inclusion) In(collection ...string) *inclusion {
	i.in = []string{}
	i.in = append(i.in, collection...)
	return i
}
