package main

import (
	"fmt"
	"strconv"
)

type day13State struct {
	x int32
	y int32
}

func (s day13State) getKey() uint64 {
	return uint64(s.x) << 32 | uint64(s.y) 
}

// Day13 link https://adventofcode.com/2016/day/13
type Day13 struct {
	favoriteNumber int
}

func (d Day13) nextStates(state day13State) []day13State {
	states := []day13State {}

	if ! d.isWall(state.x+1, state.y) { states = append(states, day13State { x: state.x+1, y: state.y }) }
	if ! d.isWall(state.x-1, state.y) { states = append(states, day13State { x: state.x-1, y: state.y }) }
	if ! d.isWall(state.x, state.y+1) { states = append(states, day13State { x: state.x, y: state.y+1 }) }
	if ! d.isWall(state.x, state.y-1) { states = append(states, day13State { x: state.x, y: state.y-1 }) }

	return states
}

func (d Day13) isWall(x int32, y int32) bool {
	if x < 0 || y < 0 { return true }
	X := int64(x)
	Y := int64(y)

	value := X*X + 3*X + 2*X*Y + Y + Y*Y + int64(d.favoriteNumber)
	bits  := false
	for value > 0 {
		if (value & 1) != 0 { bits = !bits }
		value >>= 1
	}
	return bits 
}

func (d Day13) solve() (int, int) {
	d.favoriteNumber, _ = strconv.Atoi(readline(13))

	start   := day13State { x: 1, y: 1}
	steps   := 0
	visited := make(map[uint64]bool)
	states  := make(map[uint64]day13State)

	visited[start.getKey()] = true

	states[0] = start

	answer1, answer2 := 0, 0

	for len(states) > 0 && (answer1 == 0 || answer2 == 0) {
		steps++
		if steps == 51 { 
			answer2 = len(visited)
			if answer1 != 0 {
				break 
			}
		}
		nextStates := make(map[uint64]day13State)
		for _, state := range states {
			for _,nextState := range d.nextStates(state) {
				if answer1 == 0 && nextState.x == 31 && nextState.y == 39 { 
					answer1 = steps
					if answer2 != 0 {
						break
					}
				}

				key := nextState.getKey()
				if _, ok := visited[key]; !ok {
					nextStates[key] = nextState
					visited[key] = true
				} 
			}
		}
		states = nextStates
	}

	return answer1, answer2
}


func (d Day13) run() {
	fmt.Println()
	fmt.Printf("--- Day 13 ---\n")

	part1, part2 := d.solve()

	fmt.Printf("Answer to day 13 part 1 is %v\n", part1)
	fmt.Printf("Answer to day 13 part 2 is %v\n", part2)
}
