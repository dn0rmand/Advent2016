package main

import (
	"fmt"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	} 
	return x
}

type Day1 struct {
}

func (d Day1) part1() int {
	var x, y, dx, dy = 0, 0, 0, -1

	for item := range readItems(1, ",") {
		var value, _ = strconv.Atoi(item[1:])
		switch item[0] {
			case 'R':
				dy, dx = dx, -dy
			case 'L':
				dx, dy = dy, -dx
		}
		x += dx * value
		y += dy * value
	}

	return abs(x) + abs(y)
}

func (d Day1) part2() int {
	x, y, dx, dy := 0, 0, 0, -1

	visited := make(map[int]map[int]int)

	for item := range readItems(1, ",") {
		value, _ := strconv.Atoi(item[1:])
		switch item[0] {
			case 'R':
				dy, dx = dx, -dy
			case 'L':
				dx, dy = dy, -dx
		}
		found := false
		for i := 0 ; i < value ; i++ {
			x, y += dx, dy
			y += dy
			yelem, ok := visited[y]
			if !ok { 
				yelem = make(map[int]int)
				visited[y] = yelem
			}
			_, ok = yelem[x]
			if ok {
				found = true
				break
			} 
			yelem[x] = 1	
		}
		if found { break }
	}

	return abs(x) + abs(y)
}

func (d Day1) run() {
	fmt.Println()
	fmt.Println("--- Day 1 ---")
	fmt.Printf("Answer to day 1 part 1 is %v\n", d.part1())
	fmt.Printf("Answer to day 1 part 2 is %v\n", d.part2())
}
