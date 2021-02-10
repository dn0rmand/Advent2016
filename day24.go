package main

import (
	"fmt"
	"strings"
)

type day24Bytes []byte

func (bytes day24Bytes) Len() int { 
	return len(bytes)
}

func (bytes day24Bytes) Less(i int, j int) bool { 
	return bytes[i] < bytes[j]
}

func (bytes day24Bytes) Swap(i int, j int) { 
	bytes[i], bytes[j] = bytes[j], bytes[i]
}

func (bytes day24Bytes) exists(value byte) bool { 
	for _, v := range bytes {
		if v == value {
			return true
		}
	}
	return false
}

type day24State struct {
	x int
	y int
	visited int
	complete bool
}

func (s day24State) buildNext(ox int, oy int, m[]string) *day24State {
	x := s.x + ox
	y := s.y + oy

	if x < 0 || y < 0 || y >= len(m) || x >= len(m[y]) {
		return nil
	}

	b := m[y][x]

	if b == '#' {
		return nil
	}

	if s.complete {
		return &day24State{ x: x, y: y, complete: true }
	}

	visited := s.visited

	if b >= '1' && b <= '7' {
		visited |= 1 << int(b - '1')
	}

	return &day24State{ x: x, y: y, visited: visited}
}

func (s day24State) next(m []string) []day24State {
	states := []day24State {}

	if s2 := s.buildNext(1, 0, m); s2 != nil {
		states = append(states, *s2)
	} 
	if s2 := s.buildNext(0, 1, m); s2 != nil {
		states = append(states, *s2)
	}
	if s2 := s.buildNext(-1, 0, m); s2 != nil {
		states = append(states, *s2)
	}
	if s2 := s.buildNext(0, -1, m); s2 != nil {
		states = append(states, *s2)
	}

	return states
}

func (s day24State) getKey() int {
	key := (s.x * 100 + s.y) << 8
	if s.complete {
		key |= 0xFF
	} else {
		key |= s.visited
	}
	return key
}

// Day24 link https://adventofcode.com/2016/day/24
type Day24 struct {
}

func (d Day24) loadMap() ([]string, int, int, int) {
	m := []string {}
	positions := 0
	x := 0
	y := 0

	for line := range readlines(24) {		
		m = append(m, line)
		xx := strings.IndexByte(line, '0')
		if xx >= 0 {
			x = xx
			y = len(m)-1
		}

		for _, b := range line {
			if b >= '1' && b <= '7' {
				bits := 1 << int(b - '1')
				positions |= bits
			} else if b != '.' && b != '#' && b != '0' {
					panic("only characters #.01234567 are supported")
			}
		}
	}

	return m, x, y, positions
}

func (d Day24) solve() (int, int) {
	m, x0, y0, target := d.loadMap()

	steps := 0
	part1 := -1
	part2 := -1

	visited := make(map[int]bool)
	states  := make(map[int]day24State)

	states[0] = day24State { x: x0, y: y0 }

	maxStates := 0

	for len(states) > 0 {
		if len(states) > maxStates { maxStates = len(states) }

		steps++
		newStates := make(map[int]day24State)
		for _, state := range states {
			for _, newState := range state.next(m) {
				if newState.complete && newState.x == x0 && newState.y == y0 {
					part2 = steps
					return part1, part2
				}

				if newState.visited == target {
					newState.complete = true
					if part1 < 0 { 
						part1 = steps 
					}
				}

				key := newState.getKey()
				if _, ok := visited[key]; !ok {
					newStates[key] = newState
					visited[key]   = true
				}
			}
		}

		states = newStates
	}

	return -1, -1
}

func (d Day24) run() {
	fmt.Println()
	fmt.Printf("--- Day 24 ---\n")
	part1, part2 := d.solve()

	fmt.Printf("Answer to day 24 part 1 is %v\n", part1)
	fmt.Printf("Answer to day 24 part 2 is %v\n", part2)
}
