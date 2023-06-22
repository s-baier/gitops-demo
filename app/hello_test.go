package main

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg := Hello("Gladys")
	if !want.MatchString(msg) {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, nil, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg := Hello("")
	if msg != "Hello, stranger! Are you lost?" {
		t.Fatalf(`Hello("") = %q, %v, want "Hello, stranger! Are you lost?", error`, msg, nil)
	}
}
