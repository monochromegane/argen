package ar

type Rule interface {
	Rule() *Validation
}

func MakeRule() *Validation {
	return &Validation{}
}

type Validation struct {
	presence     *presence
	length       *length
	format       *format
	numericality *numericality
	inclusion    []string
	exclusion    []string
}

func (v *Validation) Rule() *Validation {
	return v
}

func (v *Validation) Presence() *presence {
	if v.presence == nil {
		v.presence = newPresence(v)
	}
	return v.presence
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
		v.numericality = newNumericality(v)
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

type presence struct {
	*Validation
	presence bool
	message  string
}

func newPresence(v *Validation) *presence {
	return &presence{
		Validation: v,
		presence:   true,
		message:    "can't be blank",
	}
}

func (p *presence) Rule() *Validation {
	return p.Validation
}

func (p *presence) Message(message string) *presence {
	p.message = message
	return p
}

type length struct {
	*Validation
	minimum *lengthNumber
	maximum *lengthNumber
	is      *lengthNumber
}

type lengthNumber struct {
	*length
	number  int
	message string
}

func (l *lengthNumber) Rule() *Validation {
	return l.length.Validation
}

func (l *lengthNumber) Message(message string) *lengthNumber {
	l.message = message
	return l
}

func (l *length) newLengthNumber(num int, message string) *lengthNumber {
	return &lengthNumber{
		length:  l,
		number:  num,
		message: message,
	}
}

func (l *length) Rule() *Validation {
	return l.Validation
}

func (l *length) Minimum(minimum int) *lengthNumber {
	l.minimum = l.newLengthNumber(minimum, "is too short (minimum is %d characters)")
	return l.minimum
}

func (l *length) Maximum(maximum int) *lengthNumber {
	l.maximum = l.newLengthNumber(maximum, "is too long (maximum is %d characters)")
	return l.maximum
}

func (l *length) Is(is int) *lengthNumber {
	l.is = l.newLengthNumber(is, "is the wrong length (should be %d characters)")
	return l.is
}

func (l *length) In(from, to int) *lengthNumber {
	return l.Minimum(from).Maximum(to)
}

func (l *length) WithIn(from, to int) *lengthNumber {
	return l.In(from, to)
}

type format struct {
	*Validation
	with *with
}

type with struct {
	*format
	regexp  string
	message string
}

func (f *format) Rule() *Validation {
	return f.Validation
}

func (f *format) With(regexp string) *with {
	f.with = &with{
		format:  f,
		regexp:  regexp,
		message: "is invalid",
	}
	return f.with
}

func (w *with) Rule() *Validation {
	return w.format.Validation
}

func (w *with) Message(message string) *with {
	w.message = message
	return w
}

type numericality struct {
	*Validation
	numericality         bool
	message              string
	onlyInteger          *numericalityBool
	greaterThan          int
	greaterThanOrEqualTo int
	equalTo              int
	lessThan             int
	lessThanOrEqualTo    int
	odd                  *numericalityBool
	even                 *numericalityBool
}

func (n *numericality) newNumericalityBool(message string) *numericalityBool {
	return &numericalityBool{
		numericality: n,
		bool:         true,
		message:      message,
	}
}

type numericalityBool struct {
	*numericality
	bool
	message string
}

func (n *numericalityBool) Rule() *Validation {
	return n.numericality.Validation
}

func (n *numericalityBool) Message(message string) *numericalityBool {
	n.message = message
	return n
}

func newNumericality(v *Validation) *numericality {
	return &numericality{
		Validation:   v,
		numericality: true,
		message:      "is not a number",
	}
}

func (n *numericality) Rule() *Validation {
	return n.Validation
}

func (n *numericality) OnlyInteger() *numericalityBool {
	n.onlyInteger = n.newNumericalityBool("must be an integer")
	return n.onlyInteger
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

func (n *numericality) Odd() *numericalityBool {
	n.odd = n.newNumericalityBool("must be odd")
	return n.odd
}

func (n *numericality) Even() *numericalityBool {
	n.even = n.newNumericalityBool("must be even")
	return n.even
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
