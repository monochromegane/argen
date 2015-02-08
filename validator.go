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
	if v.rule.presence {
		if !v.isPersistent(value) {
			result = false
		}
	}
	if v.rule.format != nil {
		if !v.isFormatted(value) {
			result = false
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
		if !v.inLength(value) {
			result = false
		}
	}
	if v.rule.numericality != nil {
		if !v.isNumericality(value) {
			result = false
		}
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
	return result
}

func (v Validator) isPersistent(value interface{}) bool {
	return IsZero(value)
}

func (v Validator) isFormatted(value interface{}) bool {
	if v.rule.format.with == "" {
		return true
	}
	s, ok := value.(string)
	if !ok {
		return false
	}
	match, _ := regexp.MatchString(v.rule.format.with, s)
	return match
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

func (v Validator) inLength(value interface{}) bool {
	if v.rule.length.from == 0 && v.rule.length.to == 0 {
		return true
	}
	s, ok := value.(string)
	if !ok {
		return false
	}
	length := utf8.RuneCountInString(s)
	return v.rule.length.from <= length && length <= v.rule.length.to
}

func (v Validator) isNumericality(value interface{}) bool {
	result := false
	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		result = true
	case float32, float64:
		if v.rule.numericality.onlyInteger {
			result = false
		} else {
			result = true
		}
	default:
		result = false
	}
	return result
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
