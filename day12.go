package main

import (
	"fmt"
)

// Day12 link https://adventofcode.com/2016/day/12
type Day12 struct {
}

func (d Day12) part1() int {
	var a AssemBunny

	a = a.load(12)
	registers := a.run([4]int { 0, 0, 0, 0 }, nil)
	return registers[0]
}

func (d Day12) part2() int {
	var a AssemBunny

	a = a.load(12)
	registers := a.run([4]int { 0, 0, 1, 0 }, nil)
	return registers[0]
}

func (d Day12) run() {
	fmt.Println()
	fmt.Printf("--- Day 12 ---\n")
	fmt.Printf("Answer to day 12 part 1 is %v\n", d.part1())
	fmt.Printf("Answer to day 12 part 2 is %v\n", d.part2())
}
