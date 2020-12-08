package main

import (
	"testing"
)

func TestGetRow(t *testing.T) {
	var tests = []struct {
		pass string
		low  string
		high string
		want int
	}{
		{"BFFFBBF", "F", "B", 70},
		{"FFFBBBF", "F", "B", 14},
		{"BBFFBBF", "F", "B", 102},
		{"RRR", "L", "R", 7},
		{"RLL", "L", "R", 4},
	}

	for _, test := range tests {
		if convertToInt(test.pass, test.low, test.high) != test.want {
			t.Errorf("convertToInt failed. Want %v got %v", test.want, convertToInt(test.pass, test.low, test.high))
		}
	}
}
