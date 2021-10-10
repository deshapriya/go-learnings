package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Kelum"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Kelum") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

// TestHellosEmpty calls greetings.Hellos with an empty slice,
// checking for an error.
func TestHellosEmpty(t *testing.T) {
	names := []string{}
	msg, err := Hellos(names)
	if len(msg) != 0 {
		t.Fatalf(`Hellos([]) = %q, %v, want "", error`, msg, err)
	}
}
