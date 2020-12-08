package main

import (
	"testing"
)

func TestGetPosAtSlope(t *testing.T) {
	var tests = []struct {
		input [3]int
		want  int
	}{
		{[3]int{0, 1, 5}, 1},
		{[3]int{3, 3, 5}, 1},
	}

	for _, test := range tests {
		args := test.input
		if got := getNextPos(args[0], args[1], args[2]); got != test.want {
			t.Errorf("getNextPos(%q) = %v", test.input, got)
		}
	}
}
