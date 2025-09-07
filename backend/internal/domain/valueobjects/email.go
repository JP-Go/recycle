package valueobjects

import (
	"errors"
	"regexp"
)

var ErrMalformedEmail = errors.New("Malformed email")

type Email struct {
	value string
}

var baseEmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func NewEmail(emailValue string) (Email, error) {
	email := Email{value: emailValue}
	if ok := email.IsValid(); !ok {
		return Email{}, ErrMalformedEmail
	}
	return email, nil
}

func (e Email) IsEqual(o Email) bool {
	return e.value == o.value
}

func (e Email) IsValid() bool {
	return baseEmailRegex.Match([]byte(e.value))
}

func (e Email) String() string {
	return e.value
}
