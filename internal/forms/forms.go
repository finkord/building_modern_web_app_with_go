package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator/v12"
)

// Form is a struct that contains the form data and errors.
type Form struct {
	url.Values
	Errors errors
}

// Valid checks if the form is valid.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New initializes a new form.
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required checks if the form has the required fields.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field is required")
		}
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

// MinLength checks if a field has the minimum length.
func (f *Form) MinLength(field string, lenght int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < lenght {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", lenght))
		return false
	}
	return true
}

// IsEmail checks if a field is a valid email address.
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "This field must be a valid email address")
	}
}
