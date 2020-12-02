package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func product1(a []int) (c int) {
	m := make(map[int]bool)

	for _, item := range a {
		m[2020-item] = true
		if m[item] {
			fmt.Println("The two entries are:", item, 2020-item)
			c = item * (2020 - item)
			return
		}
	}
	return
}

func product2(a []int) (c int) {
	m := make(map[[2]int]bool)

	for _, item := range a {
		for _, item1 := range a {
			m[[2]int{item, 2020 - item - item1}] = true
			if m[[2]int{item, item1}] {
				fmt.Println("The three entries are:", item, item1, 2020-item-item1)
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
	fmt.Println("The two entry product is ", product1(nums))
	fmt.Println("The three entry product is ", product2(nums))
}
