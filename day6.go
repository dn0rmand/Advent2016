package main

import (
	"fmt"
	"strings"
)

// Day6 link https://adventofcode.com/2016/day/6
type Day6 struct {
}

func (d Day6) parts() (string, string) {
	var m [8][26]int

	for line := range readlines(6) {
		for i, c := range line {
			b := int(c - 'a')
			m[i][b]++
		}
	}

	var message1 strings.Builder
	var message2 strings.Builder

	for col := 0; col < 8; col++ {
		max := 0
		min := 0
		for row := 1; row < 26; row++ {
			if m[col][row] > m[col][max] {
				max = row				
			}			
			if m[col][row] < m[col][min] {
				min = row				
			}			
		}
		message1.WriteRune(rune(max+'a'))
		message2.WriteRune(rune(min+'a'))
	} 
	return message1.String(), message2.String()
}

func (d Day6) run() {
	fmt.Println()
	fmt.Printf("--- Day 6 ---\n")
	var part1, part2 = d.parts()
	fmt.Printf("Answer to day 6 part 1 is %v\n", part1)
	fmt.Printf("Answer to day 6 part 2 is %v\n", part2)
}
