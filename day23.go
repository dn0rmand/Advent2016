package main

import (
	"fmt"
)

// Day23 link https://adventofcode.com/2016/day/23
type Day23 struct {
}

func (d Day23) part1() int {
	var a AssemBunny

	a = a.load(23)
	registers := a.run([4]int { 7, 0, 0, 0 }, func(ip int, registers *[4]int) int {
		if ip == 2 {
			registers[0] = registers[0] * registers[1]
			return 10
		}
		return ip
	})
	return registers[0]
}

func (d Day23) part2() int {
	var a AssemBunny

	a = a.load(23)
	registers := a.run([4]int { 12, 0, 0, 0 }, func(ip int, registers *[4]int) int {
		if ip == 2 {
			registers[0] = registers[0] * registers[1]
			return 10
		}
		return ip
	})
	return registers[0]
}

func (d Day23) run() {
	fmt.Println()
	fmt.Printf("--- Day 23 ---\n")
	fmt.Printf("Answer to day 23 part 1 is %v\n", d.part1())
	fmt.Printf("Answer to day 23 part 2 is %v\n", d.part2())
}
