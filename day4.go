package main

import (
	"fmt"
	"strings"
	"strconv"
)

// Day4 link https://adventofcode.com/2016/day/4
type Day4 struct {
}

type day4Room struct {
	name string
	sector int
	checksum string
}

func (r day4Room) getMax(m map[rune]int) rune {
	maxRune := ' '
	maxValue:= -1
	for c, v := range m {
		if v > maxValue {
			maxValue = v
			maxRune  = c
		} else if v == maxValue && c < maxRune {
			maxRune = c
		}
	}
	return maxRune
}

func (r day4Room) isValid() bool {
	m := make(map[rune]int)

	for i := 'a'; i <= 'z'; i++ { m[i] = 0 }

	for _, c := range r.name { 
		if c != ' ' { m[c]++ }
	}

	for _, c := range r.checksum {
		if c != r.getMax(m) { return false }
		// mark as processed
		m[c] = 0
	}
	return true
}

func (r day4Room) decrypt() string {
	const A int = int('a') 

	var name strings.Builder
	name.Grow(65)

	for _, c := range r.name {
		if c == ' ' { 
			name.WriteRune(c)
		} else {
			c = rune(((int(c) - A + r.sector) % 26) + A)
			name.WriteRune(c)
		}
	}
	return name.String()
}

func (d Day4) parse(line string) day4Room {
	names := strings.Split(line, "-")
	extra := names[len(names)-1]
	names = names[0:len(names)-1]

	extras := strings.Split(extra[0:len(extra)-1], "[")

	checksum := extras[len(extras)-1]
	sector,_ := strconv.Atoi(extras[0])
	name     := strings.Join(names, " ")

	return day4Room { name, sector, checksum }
}

func (d Day4) getRooms() <-chan day4Room {
	channel := make(chan day4Room)

	go func() {
		for line := range readlines(4) {
			var r = d.parse(line)
			if r.isValid() { channel <- r }
		}
	
		close(channel)
	}()

	return channel
}

func (d Day4) part1() int {
	sum := 0

	for r := range d.getRooms() {
		sum += r.sector
	}

	return sum
}

func (d Day4) part2() int {
	for r := range d.getRooms() {
		if strings.HasPrefix(r.decrypt(), "northpole") {
			return r.sector
		}
	}
	return 0
}

func (d Day4) run() {
	fmt.Println()
	fmt.Printf("--- Day 4 ---\n")
	fmt.Printf("Answer to day 4 part 1 is %v\n", d.part1())
	fmt.Printf("Answer to day 4 part 2 is %v\n", d.part2())
}
