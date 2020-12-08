package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplitPara(t *testing.T) {
	var tests = []struct {
		input string
		want  []string
	}{
		{"a\n\nb", []string{"a", "b"}},
		{"a\nb\n\nc", []string{"a\nb", "c"}},
	}

	for _, test := range tests {
		if got := splitPara(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("getNextPos(%q) = %v", test.input, got)
		}
	}
}

func TestIsValid(t *testing.T) {
	var tests = []struct {
		pass map[string]string
		req  []string
		want bool
	}{
		{map[string]string{
			"eyr": "1972",
			"cid": "100",
			"hcl": "#18171d", "ecl": "amb", "hgt": "170", "pid": "186cm", "iyr": "2018", "byr": "1926",
		}, []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}, false},
		{map[string]string{
			"iyr": "2019",
			"hcl": "#602927", "eyr": "1967", "hgt": "170cm",
			"ecl": "grn", "pid": "012533040", "byr": "1946",
		}, []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}, false},
		{map[string]string{
			"hcl": "dab227", "iyr": "2012",
			"ecl": "brn", "hgt": "182cm", "pid": "021572410", "eyr": "2020", "byr": "1992", "cid": "277",
		}, []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}, false},
		{map[string]string{
			"hgt": "59cm", "ecl": "zzz",
			"eyr": "2038", "hcl": "74454a", "iyr": "2023",
			"pid": "3556412378", "byr": "2007",
		}, []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}, false},
		// True passports
		{map[string]string{
			"pid": "087499704", "hgt": "74in", "ecl": "grn", "iyr": "2012", "eyr": "2030", "byr": "1980",
			"hcl": "#623a2f",
		}, []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}, true},
		{map[string]string{
			"eyr": "2029", "ecl": "blu", "cid": "129", "byr": "1989",
			"iyr": "2014", "pid": "896056539", "hcl": "#a97842", "hgt": "165cm",
		}, []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}, true},
		{map[string]string{
			"hcl": "#888785",
			"hgt": "164cm", "byr": "2001", "iyr": "2015", "cid": "88",
			"pid": "545766238", "ecl": "hzl",
			"eyr": "2022",
		}, []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}, true},
		{map[string]string{
			"iyr": "2010", "hgt": "158cm", "hcl": "#b6652a", "ecl": "blu", "byr": "1944", "eyr": "2021", "pid": "093154719",
		}, []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}, true},
	}

	for _, test := range tests {
		if isValid(test.pass, test.req) != test.want {
			t.Errorf("isValid failed for input %v", test)
		}
	}
}

func testIsValidField(t *testing.T) {
	tests := []struct {
		field string
		value string
		want  bool
	}{
		{field: "byr", value: "2002", want: true},
		{field: "byr", value: "2003", want: false},
		{field: "hgt", value: "60in", want: true},
		{field: "hgt", value: "190cm", want: true},
		{field: "hgt", value: "190in", want: false},
		{field: "hgt", value: "190", want: false},
		{field: "hcl", value: "#123abc", want: true},
		{field: "hcl", value: "#123abz", want: false},
		{field: "hcl", value: "123abc", want: false},
		{field: "ecl", value: "brn", want: true},
		{field: "ecl", value: "wat", want: false},
		{field: "pid", value: "000000001", want: true},
		{field: "pid", value: "0123456789", want: false},
	}

	for _, test := range tests {
		if isValidField(test.field, test.value) != test.want {
			fmt.Println("Test failed for input:", test)
		}
	}

}
