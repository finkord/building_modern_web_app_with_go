package forms

import (
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	isValid := form.Valid()
	if !isValid {
		t.Error("form should be valid")
	}
}

func TestForm_Required(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	form = New(postedData)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when they are present")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.MinLength("a", 1)

	if form.Valid() {
		t.Error("form shows valid when min length not met")
	}

	isError := form.Errors.Get("a")
	if isError == "" {
		t.Error("form does not have error when min length not met")
	}

	postedData = url.Values{}
	postedData.Add("a", "ab")

	form = New(postedData)
	form.MinLength("a", 100)

	if form.Valid() {
		t.Error("form shows invalid when min length is met")
	}

	postedData = url.Values{}
	postedData.Add("a", "abc")

	form = New(postedData)
	form.MinLength("a", 1)

	if !form.Valid() {
		t.Error("form shows invalid when min length is met")
	}

	isError = form.Errors.Get("a")
	if isError != "" {
		t.Error("form does have error when min length met")
	}

}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	has := form.Has("a")
	if has {
		t.Error("form shows has field when it does not")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")

	form = New(postedData)
	has = form.Has("a")
	if !has {
		t.Error("form shows does not have field when it does")
	}

}

func TestForm_IsEmail(t *testing.T) {

	postedData := url.Values{}
	form := New(postedData)

	form.IsEmail("a")
	if form.Valid() {
		t.Error("form shows valid when email is not valid")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	form.IsEmail("a")
	if form.Valid() {
		t.Error("form shows valid when email is not valid")
	}

	postedData = url.Values{}
	postedData.Add("a", "a@test.com")
	form = New(postedData)

	form.IsEmail("a")
	if !form.Valid() {
		t.Error("form shows invalid when email is valid")
	}
}
