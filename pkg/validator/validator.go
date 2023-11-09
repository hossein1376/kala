package validator

import (
	"cmp"
	"regexp"
	"slices"
	"strconv"
)

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

func (v *Validator) IsNumber(value string, key, message string) {
	if _, err := strconv.Atoi(value); err != nil {
		v.AddError(key, message)
	}
}

func In[T comparable](value T, list ...T) bool {
	return slices.Contains(list, value)
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func Unique[T cmp.Ordered](values []T) bool {
	s := slices.Clone(values)
	slices.Sort(s)
	if len(slices.Compact(s)) != len(values) {
		return false
	}
	return true
}
