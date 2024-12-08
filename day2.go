package main

import (
	"fmt"
)

// Day2 link https://adventofcode.com/2016/day/2
type Day2 struct {
}

// Keypad layout
//
// 1 2 3
// 4 5 6
// 7 8 9

func (d Day2) part1() int {
	button := 4
	code   := 0

	for line := range readlines(2) {
		for _, c := range line {
			switch c {
				case 'L': if (button % 3) > 0  { button-- }
				case 'R': if (button % 3) < 2  { button++ }
				case 'U': if (button - 3 >= 0) { button -= 3 }
				case 'D': if (button + 3 <= 8) { button += 3 }
			}
		}
		code = code*10 + button+1
	}

	return code
}

// Keypad layout
//
//     1 
//   2 3 4 
// 5 6 7 8 9 
//   A B C
//	   D

// mappings
// 
//    1 2 3 4 5 6 7 8 9 A B C D
// L: 1 2 2 3 5 5 6 7 8 A A B D
// R: 1 3 4 4 6 7 8 9 9 B C C D 
// U: 1 2 1 4 5 2 3 4 9 6 7 8 B 
// D: 3 6 7 8 5 A B C 9 A D C D

func (d Day2) part2() int {

	L := [14] int { 0, 1, 2, 2, 3, 5, 5, 6, 7, 8, 10, 10, 11, 13 }
	R := [14] int { 0, 1, 3, 4, 4, 6, 7, 8, 9, 9, 11, 12, 12, 13 }
	U := [14] int { 0, 1, 2, 1, 4, 5, 2, 3, 4, 9, 6, 7, 8, 11 }
	D := [14] int { 0, 3, 6, 7, 8, 5, 10, 11, 12, 9, 10, 13, 12, 13 }

	button := 5
	code   := 0

	for line := range readlines(2) {
		for _, c := range line {
			switch c {
				case 'L': button = L[button]
				case 'R': button = R[button]
				case 'U': button = U[button]
				case 'D': button = D[button]
			}
		}
		code = code*16 + button
	}

	return code
}

func (d Day2) run() {
	fmt.Println()
	fmt.Println("--- Day 2 ---")
	fmt.Printf("Answer to day 2 part 1 is %v\n", d.part1())
	fmt.Printf("Answer to day 2 part 2 is %X\n", d.part2())
}
