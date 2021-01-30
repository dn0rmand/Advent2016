package main

import (
	"fmt"
	"strings"
	"strconv"
)

// Day9 link https://adventofcode.com/2016/day/9
type Day9 struct {
}

func (d Day9) getText() string { 
	var text strings.Builder
	for line := range readlines(9) {
		text.WriteString(line)
	}
	return text.String()
}

func (d Day9) decompress(text string, recursive bool) int {
	length := len(text)

	count := 0
	for i := 0; i < length; i++ {
		if text[i] == '(' {
			start := i + 1
			for text[i] != ')' {
				i++
			}
			end := i
			exp := strings.Split(text[start:end], "x")
			l, _ := strconv.Atoi(exp[0])
			r, _ := strconv.Atoi(exp[1])

			if recursive {
				innerText := text[i+1:i+1+l]
				l2 := d.decompress(innerText, true)
				count += l2*r
			} else {
				count += l*r
			}
			i += l
		} else {
			count++
		}
	}
	return count
}

func (d Day9) part1() int {
	return d.decompress(d.getText(), false)
}

func (d Day9) part2() int {
	return d.decompress(d.getText(), true)
}

func (d Day9) run() {
	fmt.Println()
	fmt.Printf("--- Day 9 ---\n")
	fmt.Printf("Answer to day 9 part 1 is %v\n", d.part1())
	fmt.Printf("Answer to day 9 part 2 is %v\n", d.part2())
}
