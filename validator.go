package ar

import (
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
		if !v.isMinimumLength(value) {
			result = false
		}
		if !v.isMaximumLength(value) {
			result = false
		}
		if !v.isLength(value) {
			result = false
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

func (v Validator) isMinimumLength(value interface{}) bool {
	if v.rule.length.minimum == 0 {
		return true
	}
	s, ok := value.(string)
	if !ok {
		return false
	}
	return utf8.RuneCountInString(s) <= v.rule.length.minimum
}

func (v Validator) isMaximumLength(value interface{}) bool {
	if v.rule.length.maximum == 0 {
		return true
	}
	s, ok := value.(string)
	if !ok {
		return false
	}
	return utf8.RuneCountInString(s) >= v.rule.length.maximum
}

func (v Validator) isLength(value interface{}) bool {
	if v.rule.length.is == 0 {
		return true
	}
	s, ok := value.(string)
	if !ok {
		return false
	}
	return utf8.RuneCountInString(s) == v.rule.length.is
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
