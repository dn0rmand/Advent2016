package main

import (
	"fmt"
	"strconv"
	"strings"
)

type day15Disc struct {
	start int
	positions int
}

type day15Discs []*day15Disc

// Day15 link https://adventofcode.com/2016/day/15
type Day15 struct {
}

func (d Day15) getDiscs() day15Discs {
	discs := day15Discs {}

	for line := range readlines(15) {
		number, _ := strconv.Atoi(line[6:7])
		values := strings.Split(line[12:len(line)-1], " positions; at time=0, it is at position ");
		positions, _ := strconv.Atoi(values[0])
		start, _ := strconv.Atoi(values[1])

		discs = append(discs, &day15Disc{ 
			positions: positions, 
			start: (start+number) % positions, 
		})
	}

	return discs
}

func (d Day15) calculate(discs day15Discs) int {	
	for time := 0; ; time++ {
		good := true
		for _, d := range discs {
			if d.start != 0 { good = false }
			d.start = (d.start+1) % d.positions
		}
		if good { 
			return time 
		}
	}
}

func (d Day15) part1() int {
	discs := d.getDiscs()

	return d.calculate(discs)
}

func (d Day15) part2() int {
	discs := d.getDiscs()
	discs  = append(discs, &day15Disc { positions: 11, start: len(discs)+1 })

	return d.calculate(discs)
}

func (d Day15) run() {
	fmt.Println()
	fmt.Printf("--- Day 15 ---\n")
	fmt.Printf("Answer to day 15 part 1 is %v\n", d.part1())
	fmt.Printf("Answer to day 15 part 2 is %v\n", d.part2())
}
