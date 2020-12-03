package main

import (
	"reflect"
	"testing"
)

func TestGetMatch(t *testing.T) {
	var l = "10-15 f: fluffy"
	var regexp = `(?P<min>\d+)-(?P<max>\d+) (?P<ch>\S): (?P<pass>\S+)`
	got := getMatches(regexp, l)
	want := map[string]string{
		"min":  "10",
		"max":  "15",
		"ch":   "f",
		"pass": "fluffy",
	}

	eq := reflect.DeepEqual(got, want)
	if !eq {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestIsValid(t *testing.T) {
	password := map[string]string{
		"min":  "3",
		"max":  "5",
		"ch":   "f",
		"pass": "fluffy",
	}
	got := isValidPassword(password)
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	password = map[string]string{
		"min":  "5",
		"max":  "6",
		"ch":   "f",
		"pass": "fluffy",
	}
	got = isValidPassword(password)
	want = false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestIsValidByPos(t *testing.T) {
	password := map[string]string{
		"min":  "1",
		"max":  "3",
		"ch":   "c",
		"pass": "abcde",
	}
	got := isValidPasswordByPos(password)
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	password = map[string]string{
		"min":  "1",
		"max":  "3",
		"ch":   "b",
		"pass": "cdefg",
	}
	got = isValidPasswordByPos(password)
	want = false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	password = map[string]string{
		"min":  "2",
		"max":  "9",
		"ch":   "c",
		"pass": "ccccccccc",
	}
	got = isValidPasswordByPos(password)
	want = false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
func TestGetCount(t *testing.T) {

	password := map[string]string{
		"ch":   "f",
		"pass": "fluffy",
	}
	got := countChar(password)
	want := 3

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	password = map[string]string{
		"ch":   "a",
		"pass": "banana",
	}
	got = countChar(password)
	want = 3

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	password = map[string]string{
		"ch":   "c",
		"pass": "banana",
	}
	got = countChar(password)
	want = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
