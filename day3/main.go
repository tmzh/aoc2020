package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

func getNextPos(curIdx, right, maxWidth int) int {
	return (curIdx + right) % maxWidth
}

func countTrees(lines []string, right, down int) int {
	var count, pos int
	maxWidth := len(lines[0])
	for i := 0; i < len(lines); i += down {
		if string(lines[i][pos]) == "#" {
			count++
		}
		pos = getNextPos(pos, right, maxWidth)
	}
	return count
}

func main() {
	lines := readFile("input.txt")
	count := countTrees(lines, 3, 1)
	fmt.Println("Count of trees in slope 3, 1 is :", count)

	product := 1
	var slopes = []struct {
		right int
		down  int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	for _, slope := range slopes {
		product *= countTrees(lines, slope.right, slope.down)
	}
	fmt.Println("Product of trees in slopes:", product)
}
