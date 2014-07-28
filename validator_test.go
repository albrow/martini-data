package data

import (
	"regexp"
	"testing"
)

func TestRequire(t *testing.T) {
	data := Data(map[string]string{
		"name":  "Bob",
		"age":   "25",
		"color": "",
	})
	val := data.Validator()
	val.Require("name")
	val.Require("age")
	if val.HasErrors() {
		t.Errorf("Expected no errors but got errors: %v", val.Errors)
	}

	val.Require("color")
	val.Require("a")
	if len(val.Errors) != 2 {
		t.Errorf("Expected 2 validation errors but got %d.", len(val.Errors))
	}
}

func TestMinLength(t *testing.T) {
	data := Data(map[string]string{
		"one":   "A",
		"three": "ABC",
		"five":  "ABC",
	})
	val := data.Validator()
	val.MinLength("one", 1)
	val.MinLength("three", 3)
	if val.HasErrors() {
		t.Error("Expected no errors but got errors: %v", val.Errors)
	}

	val.MinLength("five", 5)
	if len(val.Errors) != 1 {
		t.Error("Expected a validation error.")
	}
}

func TestMaxLength(t *testing.T) {
	data := Data(map[string]string{
		"one":   "A",
		"three": "ABC",
		"five":  "ABCDEF",
	})
	val := data.Validator()
	val.MaxLength("one", 1)
	val.MaxLength("three", 3)
	if val.HasErrors() {
		t.Errorf("Expected no errors but got errors: %v", val.Errors)
	}

	val.MaxLength("five", 5)
	if len(val.Errors) != 1 {
		t.Error("Expected a validation error.")
	}
}

func TestLengthRange(t *testing.T) {
	data := Data(map[string]string{
		"one-two":    "a",
		"two-three":  "abc",
		"three-four": "ab",
		"four-five":  "abcdef",
	})
	val := data.Validator()
	val.LengthRange("one-two", 1, 2)
	val.LengthRange("two-three", 2, 3)
	if val.HasErrors() {
		t.Errorf("Expected no errors but got errors: %v", val.Errors)
	}

	val.LengthRange("three-four", 3, 4)
	val.LengthRange("four-five", 4, 5)
	if len(val.Errors) != 2 {
		t.Errorf("Expected 2 validation errors but got %d.", len(val.Errors))
	}
}

func TestMatch(t *testing.T) {
	data := Data(map[string]string{
		"numeric":     "123",
		"alpha":       "abc",
		"not-numeric": "123a",
		"not-alpha":   "abc1",
	})
	val := data.Validator()
	numericRegex := regexp.MustCompile("^[0-9]+$")
	alphaRegex := regexp.MustCompile("^[a-zA-Z]+$")
	val.Match("numeric", numericRegex)
	val.Match("alpha", alphaRegex)
	if val.HasErrors() {
		t.Errorf("Expected no errors but got errors: %v", val.Errors)
	}

	val.Match("not-numeric", numericRegex)
	val.Match("not-alpha", alphaRegex)
	if len(val.Errors) != 2 {
		t.Errorf("Expected 2 validation errors but got %d.", len(val.Errors))
	}
}

func TestEmail(t *testing.T) {
	data := Data(map[string]string{
		"email":     "abc@example.com",
		"not-email": "abc.com",
	})
	val := data.Validator()
	val.Email("email")
	if val.HasErrors() {
		t.Errorf("Expected no errors but got errors: %v", val.Errors)
	}

	val.Email("not-email")
	val.Email("nothing")
	if len(val.Errors) != 2 {
		t.Errorf("Expected 2 validation errors but got %d.", len(val.Errors))
	}
}
