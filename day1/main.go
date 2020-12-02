package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func complement(a []int) (c int) {
	m := make(map[[2]int]bool)

	for _, item := range a {
		for _, item1 := range a {
			m[[2]int{item, 2020 - item - item1}] = true
			if m[[2]int{item, item1}] {
				c = item * item1 * (2020 - item - item1)
				return
			}
		}
	}
	return
}

func readFile(fname string) (nums []int, err error) {
	b, err := ioutil.ReadFile(fname)

	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil

}

func main() {
	nums, err := readFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	complement := complement(nums)
	fmt.Println(complement)
}
