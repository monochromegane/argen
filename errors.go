package ar

type Errors struct {
	Messages map[string][]error
}

func (e *Errors) Set(field string, errors []error) {
	msgs := e.message()
	msgs[field] = errors
}

func (e *Errors) Add(field string, err error) {
	msgs := e.message()
	msgs[field] = append(msgs[field], err)
}

func (e *Errors) message() map[string][]error {
	if e.Messages == nil {
		e.Messages = map[string][]error{}
	}
	return e.Messages
}
