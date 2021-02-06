package main

import (
	"crypto/md5"
	"fmt"
)

type day17State struct {
	x int
	y int
	path string
}

// Day17 link https://adventofcode.com/2016/day/17
type Day17 struct {
	prefix []byte
}

func (d Day17) openStates(state day17State) (bool, bool, bool, bool) {
	hasher := md5.New()
	hasher.Write(d.prefix)
	hasher.Write([]byte(state.path))
	hash := hasher.Sum(nil)

	// up, down, left, and right

	up := (hash[0] & 0xF0) >> 4
	down := hash[0] & 0x0F
	left := (hash[1] & 0xF0) >> 4
	right := hash[1] & 0x0F

	return (state.y > 0 && up > 10), 
				 (state.y < 3 && down > 10), 
				 (state.x > 0 && left > 10), 
				 (state.x < 3 && right > 10) 
}

func (d Day17) process() (string, int) {
	states := []day17State { { x: 0, y: 0, path: "" } }

	shortestPath := ""
	longestPath := 0 
	maxLength := 0

	for len(states) > 0 {
		maxLength++

		newStates := []day17State {}
		for _, state := range states {
			up, down, left, right := d.openStates(state)
			if up {
				s := day17State { x: state.x, y: state.y-1, path: state.path + "U" }
				if s.x == 3 && s.y == 3 { 
					if (len(shortestPath) == 0) { shortestPath = s.path }
					longestPath = maxLength
				} else {
					newStates = append(newStates, s)
				}
			}
			if down {
				s := day17State { x: state.x, y: state.y+1, path: state.path + "D" }
				if s.x == 3 && s.y == 3 { 
					if (len(shortestPath) == 0) { shortestPath = s.path }
					longestPath = maxLength
				} else {
					newStates = append(newStates, s)
				}
			}
			if left {
				s := day17State { x: state.x-1, y: state.y, path: state.path + "L" }
				if s.x == 3 && s.y == 3 { 
					if (len(shortestPath) == 0) { shortestPath = s.path }
					longestPath = maxLength
				} else {
					newStates = append(newStates, s)
				}
			}
			if right {
				s := day17State { x: state.x+1, y: state.y, path: state.path + "R" }
				if s.x == 3 && s.y == 3 { 
					if (len(shortestPath) == 0) { shortestPath = s.path }
					longestPath = maxLength
				} else {
					newStates = append(newStates, s)
				}
			}
		}
		states = newStates
	}

	return shortestPath, longestPath
}

func (d Day17) run() {
	fmt.Println()
	fmt.Printf("--- Day 17 ---\n")
	d.prefix = []byte(readline(17))
	part1, part2 := d.process()

	fmt.Printf("Answer to day 17 part 1 is %v\n", part1)
	fmt.Printf("Answer to day 17 part 2 is %v\n", part2)
}
