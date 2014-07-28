package data

import (
	"fmt"
	"regexp"
)

type Validator struct {
	data   Data
	Keys   []string
	Errors []string
}

func (v *Validator) Error(key string, err string) {
	v.Keys = append(v.Keys, key)
	v.Errors = append(v.Errors, err)
}

func (v *Validator) HasErrors() bool {
	return len(v.Errors) > 0
}

func (v *Validator) Require(key string, msg ...string) {
	if val, found := v.data[key]; !found {
		v.requiredError(key, msg...)
	} else if val == "" {
		v.requiredError(key, msg...)
	}
}

func (v *Validator) requiredError(key string, msg ...string) {
	if len(msg) != 0 {
		v.Error(key, msg[0])
	} else {
		err := fmt.Sprintf("%s is required.", key)
		v.Error(key, err)
	}
}

func (v *Validator) MinLength(key string, length int, msg ...string) {
	if val, found := v.data[key]; !found && length > 0 {
		v.minLengthError(key, length, msg...)
	} else if len(val) < length {
		v.minLengthError(key, length, msg...)
	}
}

func (v *Validator) minLengthError(key string, length int, msg ...string) {
	if len(msg) != 0 {
		v.Error(key, msg[0])
	} else {
		err := fmt.Sprintf("%s must be at least %d characters long.", key, length)
		v.Error(key, err)
	}
}

func (v *Validator) MaxLength(key string, length int, msg ...string) {
	if val, found := v.data[key]; found && len(val) > length {
		v.maxLengthError(key, length, msg...)
	}
}

func (v *Validator) maxLengthError(key string, length int, msg ...string) {
	if len(msg) != 0 {
		v.Error(key, msg[0])
	} else {
		err := fmt.Sprintf("%s cannot be more than %d characters long.", key, length)
		v.Error(key, err)
	}
}

func (v *Validator) LengthRange(key string, min int, max int, msg ...string) {
	if val, found := v.data[key]; !found && min > 0 {
		v.lengthRangeError(key, min, max, msg...)
	} else if len(val) < min || len(val) > max {
		v.lengthRangeError(key, min, max, msg...)
	}
}

func (v *Validator) lengthRangeError(key string, min int, max int, msg ...string) {
	if len(msg) != 0 {
		v.Error(key, msg[0])
	} else {
		err := fmt.Sprintf("%s must be between %d and %d characters long.", key, min, max)
		v.Error(key, err)
	}
}

func (v *Validator) Match(key string, regex *regexp.Regexp, msg ...string) {
	if val, found := v.data[key]; !found && !regex.MatchString("") {
		v.matchError(key, msg...)
	} else if !regex.MatchString(val) {
		v.matchError(key, msg...)
	}
}

func (v *Validator) Email(key string, msg ...string) {
	regex := regexp.MustCompile("^[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?$")
	v.Match(key, regex, msg...)
}

func (v *Validator) matchError(key string, msg ...string) {
	if len(msg) != 0 {
		v.Error(key, msg[0])
	} else {
		err := fmt.Sprintf("%s was not formatted correctly.", key)
		v.Error(key, err)
	}
}