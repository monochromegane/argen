package ar

import "fmt"

type Errors struct {
	Messages map[string][]error
}

func (e *Errors) SetErrors(field string, errors []error) {
	msgs := e.message()
	msgs[field] = errors
}

func (e *Errors) AddError(field string, err error) {
	msgs := e.message()
	msgs[field] = append(msgs[field], err)
}

func (e *Errors) Add(field string, err string) {
	msgs := e.message()
	msgs[field] = append(msgs[field], fmt.Errorf(err))
}

func (e *Errors) message() map[string][]error {
	if e.Messages == nil {
		e.Messages = map[string][]error{}
	}
	return e.Messages
}

func (e *Errors) Error() string {
	resp := ""
	msgs := e.message()
	for key, errs := range msgs {
		resp += key + ": "
		for _, m := range errs {
			resp += m.Error() + ". "
		}
	}
	return resp
}
