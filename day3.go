package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Day3 link https://adventofcode.com/2016/day/3
type Day3 struct {
}

func (d Day3) isTriangle(v1 int, v2 int, v3 int) bool {
	return v1+v2 > v3 && v2+v3 > v1 && v3+v1 > v2
}

func (d Day3) part1() int {
	count := 0
	for line := range readlines(3) {
		v1, _ := strconv.Atoi(strings.TrimSpace(line[0:5]))
		v2, _ := strconv.Atoi(strings.TrimSpace(line[5:10]))
		v3, _ := strconv.Atoi(strings.TrimSpace(line[10:]))

		if d.isTriangle(v1, v2, v3) {
			count++
		}
	}
	return count
}

func (d Day3) part2() int {
	var count, row = 0, 0
	var v11, v12, v21, v22, v31, v32 = 0, 0, 0, 0, 0, 0

	for line := range readlines(3) {
		v13, _ := strconv.Atoi(strings.TrimSpace(line[0:5]))
		v23, _ := strconv.Atoi(strings.TrimSpace(line[5:10]))
		v33, _ := strconv.Atoi(strings.TrimSpace(line[10:]))

		row = (row + 1) % 3
		if row == 0 {
			if d.isTriangle(v11, v12, v13) { count++ }
			if d.isTriangle(v21, v22, v23) { count++ }
			if d.isTriangle(v31, v32, v33) { count++ }
		}

		v11, v12 = v12, v13
		v21, v22 = v22, v23
		v31, v32 = v32, v33
	}

	return count
}

func (d Day3) run() {
	fmt.Println()
	fmt.Printf("--- Day 3 ---\n")
	fmt.Printf("Answer to day 3 part 1 is %v\n", d.part1())
	fmt.Printf("Answer to day 3 part 2 is %v\n", d.part2())
}
