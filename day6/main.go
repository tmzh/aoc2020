package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type groupAnswers struct {
	answers []string
}

func unique(arr []string) []string {
	occurred := map[string]bool{}
	for e := range arr {
		occurred[arr[e]] = true
	}
	result := []string{}
	for keys := range occurred {
		result = append(result, keys)
	}

	return result
}

func readFile(fname string) (groups []string) {
	b, err := ioutil.ReadFile(fname)

	if err != nil {
		panic(err)
	}

	groups = splitGroup(string(b))
	return
}

func splitGroup(text string) []string {
	return strings.Split(text, "\n\n")
}

func distinctInGroup(group string) groupAnswers {
	entries := strings.Split(strings.Replace(group, "\n", "", -1), "")
	g := groupAnswers{
		answers: unique(entries),
	}
	return g
}

func intersect(s1, s2 string) string {
	var result string
	s1Map := map[rune]bool{}
	for _, k := range s1 {
		s1Map[k] = true
	}

	for _, k := range s2 {
		if s1Map[k] {
			result = result + string(k)
		}
	}

	return result

}

func intersectReduce(s []string) string {
	if len(s) == 1 {
		return s[0]
	} else if len(s) == 2 {
		return intersect(s[0], s[1])
	}
	return intersectReduce(append(s[2:], intersect(s[0], s[1])))
}

func commonInGroup(group string) groupAnswers {
	individuals := strings.Split(group, "\n")
	g := groupAnswers{
		answers: strings.Split(intersectReduce(individuals), ""),
	}
	return g
}

func main() {
	groups := readFile("input.txt")
	allAnswers := make([]groupAnswers, len(groups))
	for _, group := range groups {
		allAnswers = append(allAnswers, distinctInGroup(group))
	}
	var count int

	for _, g := range allAnswers {
		count += len(g.answers)
	}

	commonAnswers := make([]groupAnswers, len(groups))
	for _, group := range groups {
		commonAnswers = append(commonAnswers, commonInGroup(group))
	}
	var commonCount int

	for _, g := range commonAnswers {
		commonCount += len(g.answers)
	}

	fmt.Println("Number of unique answers:", count)
	fmt.Println("Number of intra group common answers:", commonCount)

}
