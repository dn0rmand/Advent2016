package main

import (
	"fmt"
	"strings"
	"strconv"
)

// Day4 link https://adventofcode.com/2016/day/4
type Day4 struct {
}

func parse(line string) ([]string, int, string) {
	names := strings.Split(line, "-")
	extra := names[len(names)-1]
	names = names[0:len(names)-1]

	extras := strings.Split(extra[0:len(extra)-1], "[")

	checksum := extras[len(extras)-1]
	sector,_ := strconv.Atoi(extras[0])

	return names, sector, checksum
}

func isValid(names []string, checksum string) bool {
	return true
}

func (d Day4) part1() int {
	sum := 0

	for line := range readlines(4) {
		var names, sector, checksum = parse(line)
		if isValid(names, checksum) { sum += sector }
	}

	return sum
}

func (d Day4) part2() int {
	// for line := range readlines(4) {
	// }
	return 0
}

func (d Day4) run() {
	fmt.Println()
	fmt.Printf("--- Day 4 ---\n")
	fmt.Printf("Answer to day 4 part 1 is %v\n", d.part1())
	fmt.Printf("Answer to day 4 part 2 is %v\n", d.part2())
}
