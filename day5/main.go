package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func readFile(fname string) (lines []string) {
	b, err := ioutil.ReadFile(fname)

	if err != nil {
		panic(err)
	}

	lines = strings.Split(string(b), "\n")
	return
}

func convertToInt(text, low, high string) int {
	binString := strings.Replace(text, low, "0", -1)
	binString = strings.Replace(binString, high, "1", -1)
	s, _ := strconv.ParseInt(binString, 2, 32)
	return int(s)
}

func findGap(rowIds []int) int {
	min := rowIds[0]
	for idx, rowID := range rowIds {
		if rowID > (idx + min) {
			return idx + min
		}
	}
	return -1
}

func main() {
	seats := readFile("input.txt")
	var highestRowID int
	var rowIds []int

	for _, seat := range seats {
		if len(seat) == 0 {
			continue
		}
		row, col := convertToInt(seat[:7], "F", "B"), convertToInt(seat[7:], "L", "R")
		rowID := row*8 + col
		rowIds = append(rowIds, rowID)
		if rowID > highestRowID {
			highestRowID = rowID
		}
	}
	fmt.Printf("The larged row id is %v\n", highestRowID)
	sort.Ints(rowIds)
	fmt.Printf("The missing row id is %v\n", findGap(rowIds))
}
