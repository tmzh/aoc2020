package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func readFile(fname string) (paras []string) {
	b, err := ioutil.ReadFile(fname)

	if err != nil {
		panic(err)
	}

	paras = splitPara(string(b))
	return
}

func splitPara(text string) []string {
	return strings.Split(text, "\n\n")
}

func paraToMap(para string) map[string]string {
	m := make(map[string]string)
	entries := strings.Split(strings.Replace(para, "\n", " ", -1), " ")
	for _, entry := range entries {
		if len(entry) == 0 {
			continue
		}
		kv := strings.Split(entry, ":")
		m[kv[0]] = kv[1]
	}
	return m
}

/*
   byr (Birth Year) - four digits; at least 1920 and at most 2002.
   iyr (Issue Year) - four digits; at least 2010 and at most 2020.
   eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
   hgt (Height) - a number followed by either cm or in:
       If cm, the number must be at least 150 and at most 193.
       If in, the number must be at least 59 and at most 76.
   hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
   ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
   pid (Passport ID) - a nine-digit number, including leading zeroes.
   cid (Country ID) - ignored, missing or not.
*/
func isValidField(field, value string) bool {
	switch field {
	case "byr":
		byr, _ := strconv.Atoi(value)
		return byr >= 1920 && byr <= 2002
	case "iyr":
		iyr, _ := strconv.Atoi(value)
		return iyr >= 2010 && iyr <= 2020
	case "eyr":
		eyr, _ := strconv.Atoi(value)
		return eyr >= 2020 && eyr <= 2030
	case "hgt":
		if strings.HasSuffix(value, "cm") {
			return value >= "150cm" && value <= "193cm"
		} else if strings.HasSuffix(value, "in") {
			return value >= "59in" && value <= "76in"
		} else {
			return false
		}
	case "hcl":
		re := regexp.MustCompile(`^#[0-9a-f]{6}`)
		return re.Match([]byte(value))
	case "ecl":
		re := regexp.MustCompile(`(amb|blu|brn|gry|grn|hzl|oth)`)
		return re.Match([]byte(value))
	case "pid":
		re := regexp.MustCompile(`[0-9]{9}`)
		return re.Match([]byte(value))
	default:
		return true
	}
}

func isValid(pass map[string]string, req []string) bool {
	for _, field := range req {
		if value, ok := pass[field]; !ok {
			return false
		} else if !isValidField(field, value) {
			return false
		}
	}
	return true
}

func main() {
	paras := readFile("input.txt")
	credsMap := make([]map[string]string, 0, len(paras))
	for _, para := range paras {
		credsMap = append(credsMap, paraToMap(para))
	}

	fakePass := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	var count int

	for _, pass := range credsMap {
		if isValid(pass, fakePass) {
			count++
		}
	}

	fmt.Println("Number of fake passports:", count)

}
