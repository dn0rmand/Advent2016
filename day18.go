package main

import (
	"fmt"
)

// Day18 link https://adventofcode.com/2016/day/18
type Day18 struct {
}

func (d Day18) countSafeTiles(rows int) int {
	row := []byte(readline(18))
	l   := len(row)

	count := 0
	for _, c := range row { 
		if c == '.' { count++ }
	}

	for i := 1; i < rows; i++ {
		newRow := make([]byte, l)
	
		count += l
		for c := 0; c < l; c++ {
			newRow[c] = '.'
			left := byte('.')
			right:= byte('.')
			if c > 0 { left = row[c-1] }
			if c < l-1 { right = row[c+1] }

			if left == '^' && right == '.' { 
				newRow[c] = '^' 
				count--
			} else if left == '.' && right == '^' {
				newRow[c] = '^'
				count--
			} 
		}

		row = newRow
	}

	return count
}

func (d Day18) part1() int {
	return d.countSafeTiles(40)
}

func (d Day18) part2() int {
	return d.countSafeTiles(400000)
}

func (d Day18) run() {
	fmt.Println()
	fmt.Printf("--- Day 18 ---\n")
	fmt.Printf("Answer to day 18 part 1 is %v\n", d.part1())
	fmt.Printf("Answer to day 18 part 2 is %v\n", d.part2())
}
