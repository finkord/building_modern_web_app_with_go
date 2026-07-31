package forms

import (
	"net/http"
	"net/url"
)

// Form is a struct that contains the form data and errors.
type Form struct {
	url.Values
	Errors errors
}

// New initializes a new form.
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if a field exists in the form.
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}

	return true
}
