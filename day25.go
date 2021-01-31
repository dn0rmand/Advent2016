package main

import (
	"fmt"
)

// Day25 link https://adventofcode.com/2016/day/25
type Day25 struct {
}

func (d Day25) part1() int {
	var a AssemBunny

	next := 0
	count:= 0

	a.output = func(value int) bool { 
		if next != value {
			return true
		}
		next = (next+1) % 2
		count++
		return count > 100
	}

	a = a.load(25)

	for regA := 1; ; regA++ { 
		next = 0
		count = 0
		a.run([4]int { regA, 0, 0, 0 }, nil)
		if count > 100 { 
			return regA 
		}
	}
}

func (d Day25) run() {
	fmt.Println()
	fmt.Printf("--- Day 25 ---\n")
	fmt.Printf("Answer to day 25 part 1 is %v\n", d.part1())
}
