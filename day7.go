package main

import (
	"fmt"
)

// Day7 link https://adventofcode.com/2016/day/7
type Day7 struct {
}

func (d Day7) isABBA(ip string) bool {
	inside := false

	l := len(ip)

	good := false

	for i, c := range ip {		
		if i+4 > l { break }

		if c == '[' { 
			inside = true 
		} else if c == ']' { 
			inside = false 
		} else {
			if ip[i+1] == ']' || ip[i+1] == '[' { continue }
			if ip[i] == ip[i+3] && ip[i+1] == ip[i+2] && ip[i] != ip[i+1] {
				if inside { return false }
				good = true
			}
		}
	}

	return good
}

func (d Day7) isABA(ip string) bool {
	inside := false

	ABA := make(map[string]string)
	BAB := make(map[string]string)

	l := len(ip)

	for i, c := range ip {		
		if i+3 > l { break }

		if c == '[' { 
			inside = true 
		} else if c == ']' { 
			inside = false 
		} else {
			if ip[i+1] == ']' || ip[i+1] == '[' { continue }
			if ip[i] == ip[i+2] && ip[i] != ip[i+1] {
				k1 := fmt.Sprintf("%c%c%c", c, ip[i+1], c)
				k2 := fmt.Sprintf("%c%c%c", ip[i+1], c, ip[i+1])
				if inside { 
					BAB[k1] = k2
				} else {
					ABA[k1] = k2
				}
			}
		}
	}

	for k1, k2 := range ABA {
		if BAB[k2] == k1 { return true }
	}
	return false
}

func (d Day7) part1() int {
	count := 0

	for line := range readlines(7) {
		if d.isABBA(line) { count++ }
	}
	return count
}

func (d Day7) part2() int {
	count := 0

	for line := range readlines(7) {
		if d.isABA(line) { count++ }
	}
	return count
}

func (d Day7) run() {
	fmt.Println()
	fmt.Printf("--- Day 7 ---\n")
	fmt.Printf("Answer to day 7 part 1 is %v\n", d.part1())
	fmt.Printf("Answer to day 7 part 2 is %v\n", d.part2())
}
