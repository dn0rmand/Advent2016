package main

import (
	"fmt"
	"strconv"
	"strings"
)

type day22Disk struct {
	used int
	free int
}

type day22State struct {
	maxX int
	maxY int
	values map[int]day22Disk
}

func (s day22State) findEmpty() (int, int) {
	id := -1 
	for k, v := range s.values { 
		if v.used == 0 {
			if id == -1 { id = k } else { panic("More than 1 empty disk") }
		}
	}

	y := id % 100
	x := (id - y)/100
	return x, y
}

// Day22 link https://adventofcode.com/2016/day/22
type Day22 struct {
}

func (d Day22) load() day22State {
	maxX := 0
	maxY := 0
	data := make(map[int]day22Disk)

	for line := range readlines(22) {
		if strings.HasPrefix(line, "/dev/grid/node-x") {
			XY     := strings.Split(line[15:22], "-")
			used,_ := strconv.Atoi(strings.Trim(line[30:33], " "))
			free,_ := strconv.Atoi(strings.Trim(line[37:40], " "))
			x,_    := strconv.Atoi(strings.Trim(XY[0][1:], " "))
			y,_    := strconv.Atoi(strings.Trim(XY[1][1:], " "))

			k := x*100 + y
			data[k] = day22Disk{ used: used, free: free }
			if x > maxX { maxX = x }
			if y > maxY { maxY = y }
		}
	}

	return day22State{
		maxX: maxX,
		maxY: maxY,
		values: data,
	}
}

func (d Day22) part1() int {
	values := d.load().values

	viable := 0

	for _, disk1 := range values {
		if disk1.used == 0 { continue }
		for _, disk2 := range values {
			if disk1 == disk2 { continue }
			if disk1.used <= disk2.free {
				viable++
			}
		}
	}

	return viable
}

func (d Day22) part2() int {
	info := d.load()

	x, y := info.findEmpty()

	states := make(map[int]int)
	states[x*100 + y] = 1

	steps  := -1
	for len(states) > 0 {		
		steps++
		nextStates := make(map[int]int)
		for key := range states {
			y := key % 100
			x := (key - y) / 100

			if (x == info.maxX-1 && y == 0) { 
				nextStates = make(map[int]int)
				break
			}
			if (x == info.maxX && y == 1){ 
				nextStates = make(map[int]int)
				break
			}

			disk1 := info.values[key]
			size := disk1.used + disk1.free

			if (x > 0) {
				disk2 := info.values[key-100]
				if (disk2.used <= size) {
					nextStates[key-100] = 1
				}
			}
			if (x < info.maxX) {
				disk2 := info.values[key+100]
				if (disk2.used <= size) {
					nextStates[key+100] = 1
				}
			}
			if (y > 0) {
				disk2 := info.values[key-1]
				if (disk2.used <= size) {
					nextStates[key-1] = 1
				}
			}
			if (y < info.maxY) {
				disk2 := info.values[key+1]
				if (disk2.used <= size) {
					nextStates[key+1] = 1
				}
			}
		}
		states = nextStates
	}

	steps = steps + 1 + ((info.maxX-1) * 5)
	return steps
}

func (d Day22) run() {
	fmt.Println()
	fmt.Printf("--- Day 22 ---\n")
	fmt.Printf("Answer to day 22 part 1 is %v\n", d.part1())
	fmt.Printf("Answer to day 22 part 2 is %v\n", d.part2())
}
