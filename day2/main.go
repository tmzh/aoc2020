package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func countChar(matchMap map[string]string) (count int) {
	return strings.Count(matchMap["pass"], matchMap["ch"])
}

func isValidPassword(matchMap map[string]string) bool {
	max, err := strconv.Atoi(matchMap["max"])
	if err != nil {
		panic(err)
	}

	min, err := strconv.Atoi(matchMap["min"])
	if err != nil {
		panic(err)
	}
	return (min <= countChar(matchMap)) && (countChar(matchMap) <= max)
}

func isValidPasswordByPos(matchMap map[string]string) bool {
	max, err := strconv.Atoi(matchMap["max"])
	if err != nil {
		panic(err)
	}

	min, err := strconv.Atoi(matchMap["min"])
	if err != nil {
		panic(err)
	}
	pass := matchMap["pass"]
	ch := matchMap["ch"]

	var matches int

	if string(pass[max-1]) == ch {
		matches++
	}

	if string(pass[min-1]) == ch {
		matches++
	}

	return matches == 1
}

func getMatches(regEx, url string) (matchMap map[string]string) {

	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(url)

	matchMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			matchMap[name] = match[i]
		}
	}
	return
}

func readFile(fname string) []string {
	b, err := ioutil.ReadFile(fname)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	linesList := make([]string, 0, len(lines))
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		linesList = append(linesList, l)
	}
	return linesList
}

func main() {
	lines := readFile("input.txt")
	var count int
	var posCount int
	for _, l := range lines {
		var regexp = `(?P<min>\d+)-(?P<max>\d+) (?P<ch>\S): (?P<pass>\S+)`
		matchMap := getMatches(regexp, l)
		if isValidPassword(matchMap) {
			count++
		}
		if isValidPasswordByPos(matchMap) {
			posCount++
		}
	}
	fmt.Println("Valid passwords by count", count)
	fmt.Println("Valid passwords by position", posCount)
}
