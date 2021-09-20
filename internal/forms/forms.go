package forms

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Forms struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Forms {
	return &Forms{
		data,
		errors{},
	}
}

func (f *Forms) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

func (f *Forms) MinLength(field string, length int) {
	value := f.Get(field)
	if len(value) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
	}
}

func (f *Forms) Valid() bool {
	return len(f.Errors) == 0
}

func (f *Forms) RegisteredEmail(field string, r *http.Request) {
	if field == "email" {
		value := f.Get(field)
		c, _ := r.Cookie(field)

		log.Println("cookie c", c)
		if c != nil && c.Value != "" && c.Value == value {
			f.Errors.Add("email", "The EMail is already registered")
		}
	}

}
