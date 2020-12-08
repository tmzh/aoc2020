package main

import (
	"reflect"
	"testing"
)

func TestSplitPara(t *testing.T) {
	var tests = []struct {
		input string
		want  []string
	}{
		{"abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb", []string{"abc", "a\nb\nc", "ab\nac", "a\na\na\na", "b"}},
	}

	for _, test := range tests {
		if got := splitGroup(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("splitGroup(%q) = %v", test.input, got)
		}
	}
}

func TestGetDistinctInGroup(t *testing.T) {
	var tests = []struct {
		input string
		want  []string
	}{
		{"abc", []string{"a", "b", "c"}},
		{"a\nb\nc", []string{"a", "b", "c"}},
		{"ab\nac", []string{"a", "b", "c"}},
		{"a\na\na\na", []string{"a"}},
	}

	for _, test := range tests {
		if got := distinctInGroup(test.input); !reflect.DeepEqual(got.answers, test.want) {
			t.Errorf("splitGroup(%q) = %v", test.input, got)
		}
	}
}

func TestGetIntersection(t *testing.T) {
	var tests = []struct {
		input [2]string
		want  string
	}{
		{[2]string{"ab", "bc"}, "b"},
		{[2]string{"a", "b"}, ""},
	}

	for _, test := range tests {
		if got := intersect(test.input[0], test.input[1]); !reflect.DeepEqual(got, test.want) {
			t.Errorf("splitGroup(%q) = %v", test.input, got)
		}
	}
}

func TestGetIntersectionRecursive(t *testing.T) {
	var tests = []struct {
		input []string
		want  string
	}{
		{[]string{"abc"}, "abc"},
		{[]string{"ab", "ac"}, "a"},
		{[]string{"a", "a", "a"}, "a"},
	}

	for _, test := range tests {
		if got := intersectReduce(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("splitGroup(%q) = %v", test.input, got)
		}
	}
}
