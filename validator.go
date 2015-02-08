package ar

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

type Validator struct {
	rule *Validation
}

func NewValidator(rule *Validation) Validator {
	return Validator{rule}
}

func (v Validator) IsValid(value interface{}) bool {
	result := true
	errors := []error{}
	if v.rule.presence != nil {
		if ok, err := v.isPersistent(value); !ok {
			result = false
			errors = append(errors, err)
		}
	}
	if v.rule.format != nil {
		if ok, err := v.isFormatted(value); !ok {
			result = false
			errors = append(errors, err)
		}
	}
	if v.rule.length != nil {
		if ok, err := v.isMinimumLength(value); !ok {
			result = false
			errors = append(errors, err)
		}
		if ok, err := v.isMaximumLength(value); !ok {
			result = false
			errors = append(errors, err)
		}
		if ok, err := v.isLength(value); !ok {
			result = false
			errors = append(errors, err)
		}
		if ok, err := v.inLength(value); !ok {
			result = false
			errors = append(errors, err)
		}
	}
	if v.rule.numericality != nil {
		if ok, err := v.isNumericality(value); !ok {
			result = false
			errors = append(errors, err)
		} else {
			if !v.greaterThan(value) {
				result = false
			}
			if !v.greaterThanOrEqualTo(value) {
				result = false
			}
			if !v.equalTo(value) {
				result = false
			}
			if !v.lessThan(value) {
				result = false
			}
			if !v.lessThanOrEqualTo(value) {
				result = false
			}
			if !v.odd(value) {
				result = false
			}
			if !v.even(value) {
				result = false
			}
		}
	}
	return result
}

func (v Validator) isPersistent(value interface{}) (bool, error) {
	if IsZero(value) {
		return false, fmt.Errorf("%s", v.rule.presence.message)
	}
	return true, nil
}

func (v Validator) isFormatted(value interface{}) (bool, error) {
	with := v.rule.format.with
	if with.regexp == "" {
		return true, nil
	}
	s, ok := value.(string)
	if !ok {
		return false, fmt.Errorf(with.message)
	}
	match, _ := regexp.MatchString(with.regexp, s)
	if !match {
		return false, fmt.Errorf(with.message)
	}
	return true, nil
}

func (v Validator) isMinimumLength(value interface{}) (bool, error) {
	minimum := v.rule.length.minimum
	if minimum.number == 0 {
		return true, nil
	}
	result := utf8.RuneCountInString(fmt.Sprintf("%s", value)) <= minimum.number
	if !result {
		return false, fmt.Errorf(minimum.message, minimum.number)
	}
	return true, nil
}

func (v Validator) isMaximumLength(value interface{}) (bool, error) {
	maximum := v.rule.length.maximum
	if maximum.number == 0 {
		return true, nil
	}
	result := utf8.RuneCountInString(fmt.Sprintf("%s", value)) >= maximum.number
	if !result {
		return false, fmt.Errorf(maximum.message, maximum.number)
	}
	return true, nil
}

func (v Validator) isLength(value interface{}) (bool, error) {
	is := v.rule.length.is
	if is.number == 0 {
		return true, nil
	}
	result := utf8.RuneCountInString(fmt.Sprintf("%s", value)) == is.number
	if !result {
		return false, fmt.Errorf(is.message, is.number)
	}
	return true, nil
}

func (v Validator) inLength(value interface{}) (bool, error) {
	ok, err := v.isMinimumLength(value)
	if !ok {
		return false, err
	}
	ok, err = v.isMaximumLength(value)
	if !ok {
		return false, err
	}
	return true, nil
}

func (v Validator) isNumericality(value interface{}) (bool, error) {
	numericality := v.rule.numericality
	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true, nil
	case float32, float64:
		if numericality.onlyInteger {
			return false, fmt.Errorf("must be an integer")
		} else {
			return true, nil
		}
	}
	return false, fmt.Errorf("is not a number")
}

func (v Validator) greaterThan(value interface{}) bool {
	i, ok := value.(int)
	if !ok {
		return false
	}
	return i > v.rule.numericality.greaterThan
}

func (v Validator) greaterThanOrEqualTo(value interface{}) bool {
	i, ok := value.(int)
	if !ok {
		return false
	}
	return i >= v.rule.numericality.greaterThanOrEqualTo
}

func (v Validator) equalTo(value interface{}) bool {
	i, ok := value.(int)
	if !ok {
		return false
	}
	return i == v.rule.numericality.equalTo
}

func (v Validator) lessThan(value interface{}) bool {
	i, ok := value.(int)
	if !ok {
		return false
	}
	return i < v.rule.numericality.lessThan
}

func (v Validator) lessThanOrEqualTo(value interface{}) bool {
	i, ok := value.(int)
	if !ok {
		return false
	}
	return i <= v.rule.numericality.lessThanOrEqualTo
}

func (v Validator) odd(value interface{}) bool {
	if !v.rule.numericality.odd {
		return true
	}
	i, ok := value.(int)
	if !ok {
		return false
	}
	return i%2 == 1
}

func (v Validator) even(value interface{}) bool {
	if !v.rule.numericality.even {
		return true
	}
	i, ok := value.(int)
	if !ok {
		return false
	}
	return i%2 == 0
}
