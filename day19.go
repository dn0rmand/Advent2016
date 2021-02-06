package main

import (
	"fmt"
	"strconv"
)

type day19Elf struct {
	id int
	next *day19Elf
	previous *day19Elf
}

// Day19 link https://adventofcode.com/2016/day/19
type Day19 struct {
}

func (d Day19) loadElves() (*day19Elf, int) {
	count, _ := strconv.Atoi(readline(19))

	first := &day19Elf{ id: 1 }
	last  := first
	
	for e := 2; e <= count; e++ {
		last.next = &day19Elf{ id: e, previous: last }
		last = last.next
	}

	last.next = first
	first.previous = last

	return first, count
}

func (d Day19) part1() int {
	elf, count := d.loadElves();

	for count > 1 {
		elf.next = elf.next.next
		elf = elf.next
		count--
	}

	return elf.id
}

func (d Day19) part2() int {
	current, count := d.loadElves();

	offset := 0
	opposite := current
	for count > 1 {
		newOffset := count >> 1

		for offset < newOffset {
			offset++
			opposite = opposite.next
		}
		for offset > newOffset {
			offset--
			opposite = opposite.previous
		}

		// remove opposite
		o1 := opposite.previous
		opposite = opposite.next
		o1.next = opposite
		opposite.previous = o1
		opposite = opposite.next

		// next elf's turn
		current = current.next
		count--
	}

	return current.id
}

func (d Day19) run() {
	fmt.Println()
	fmt.Printf("--- Day 19 ---\n")
	fmt.Printf("Answer to day 19 part 1 is %v\n", d.part1())
	fmt.Printf("Answer to day 19 part 2 is %v\n", d.part2())
}
