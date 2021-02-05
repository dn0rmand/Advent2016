package main

import (
	"fmt"
)

// Day16 link https://adventofcode.com/2016/day/16
type Day16 struct {
}

func (d Day16) fillDisk(size int) []byte {
	data := []byte {}
	for _, b := range readline(16) {
		if b == '0' { data = append(data, 0) } else { data = append(data, 1) }
	}
	for len(data) < size {
		l := len(data)
		copy := make([]byte, l)

		for i := 0; i < l; i++ { 
			if data[l-i-1] == 0 { copy[i] = 1 } else { copy[i] = 0 }
		} 
		data = append(data, 0)
		data = append(data, copy...)
	}

	return data[:size]
}

func (d Day16) checkSum(disc []byte) string {	
	for (len(disc) & 1) == 0 {
		l := len(disc) >> 1
		check := make([]byte, l)

		for i := 0; i < l; i++ {
			if disc[i*2] == disc[i*2 + 1] {
				check[i] = 1
			} else {
				check[i] = 0
			}
		}

		disc = check
	}
	for i := 0; i < len(disc); i++ {
		disc[i] += '0'
	}

	return string(disc)
}

func (d Day16) part1() string {
	disc := d.fillDisk(272)

	checksum := d.checkSum(disc)

	return checksum
}

func (d Day16) part2() string {
	disc := d.fillDisk(35651584)

	checksum := d.checkSum(disc)

	return checksum
}

func (d Day16) run() {
	fmt.Println()
	fmt.Printf("--- Day 16 ---\n")
	fmt.Printf("Answer to day 16 part 1 is %v\n", d.part1())
	fmt.Printf("Answer to day 16 part 2 is %v\n", d.part2())
}
