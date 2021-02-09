package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var day20MAX = 4294967295

type day20Ranges []*day20Range

type day20Range struct {
	min int
	max int
}

// Day20 link https://adventofcode.com/2016/day/20
type Day20 struct {
}

func (r day20Ranges) Len() int { return len(r) }
func (r day20Ranges) Swap(i int, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r day20Ranges) Less(i int, j int) bool {
	if (r[i].min == r[j].min) {
		return r[i].max < r[j].max
	}
	return r[i].min < r[j].min
}

func (d Day20) parts() (int, int) {
	ranges := day20Ranges {}

	for line := range readlines(20) {
		values := strings.Split(line,"-")
		min, _ := strconv.Atoi(values[0])
		max, _ := strconv.Atoi(values[1])

		done := false
		for _,r := range ranges {
			if max >= r.min && max <= r.max {
				if min < r.min  {
					r.min = min
					done  = true
				}
				done  = true
			} else if (min >= r.min && min <= r.max) {
				if max > r.max {
					r.max = max
				}
				done = true
			}
		}
		if ! done {
			ranges = append(ranges, &day20Range { min: min, max: max })
		}
	}

	sort.Sort(ranges)

	for i := 0; i < len(ranges)-1; i++ {
		r1 := ranges[i] 
		r2 := ranges[i+1]

		if r1.max+1 >= r2.min {
			if r1.max < r2.max {
				r1.max = r2.max
			}
			if (i+2 == len(ranges)) {
				ranges = ranges[0:i+1]
			} else {
				ranges = append(ranges[0:i+1], ranges[i+2:]...)
			}
			i--
		} 
	}

	total := day20MAX+1
	minValue := 0

	for _, r := range ranges {

		if r.min <= minValue {
			minValue = r.max+1
		}

		total -= (r.max - r.min + 1)
	}

	return minValue, total
}

func (d Day20) run() {
	fmt.Println()
	fmt.Printf("--- Day 20 ---\n")
	part1, part2 := d.parts()

	fmt.Printf("Answer to day 20 part 1 is %v\n", part1)
	fmt.Printf("Answer to day 20 part 2 is %v\n", part2)
}
