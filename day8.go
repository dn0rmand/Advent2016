package main

import (
	"fmt"
	"strconv"
	"strings"
)

// rect 1x1
// rotate row y=0 by 5
// rotate column x=0 by 1

// Day8 link https://adventofcode.com/2016/day/8
type Day8 struct {
	screen [6][50]rune
}

func (d Day8) print() {
	for _, row := range d.screen {
		for _, c := range row {
			if c != 0 { 
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (d Day8) count(letter int) int {
	count := 0
	var start, end = 0, 50

	if letter >= 0 {
		start = letter * 5
		end   = start + 5
	}

	for col := start; col < end; col++ {
		for row := 0; row < 6; row++ {
			if d.screen[row][col] != 0 { count++ }
		}
	}

	return count
}

func (d Day8) getLetter(letter int) rune {
	ocr := [26]int { 0x19297A52, // A 
									 0x392E4A5C, // B
									 0x1928424C, // C
									 0x39294A5C, // D
									 0x3D0E421E, // E
									 0x3D0E4210, // F
									 0x19285A4E, // G
									 0x252F4A52, // H
									 0x1C42108E, // I
									 0x0C210A4C, // J
									 0x254C5292, // K
									 0x2108421E, // L
									 0x00000000, // M
									 0x00000000, // N
									 0x19294A4C, // O
									 0x39297210, // P
									 0x00000000, // Q
									 0x39297292, // R
									 0x1D08305C, // S
									 0x1C421084, // T
									 0x25294A4C, // U
									 0x00000000, // V
									 0x00000000, // W
									 0x00000000, // X
									 0x23151084, // Y
									 0x3C22221E, // Z
									}

	start := letter * 5
	end   := start + 5 // actually only 4 pixels wide

	value := 0

	for row := 0; row < 6; row++ {
		for col := start; col < end; col++ {
			if d.screen[row][col] != 0 {
				value = (value << 1) | 1
			} else {
				value = value << 1
			}			
		}
	}

	A := byte('A')
	for i, v := range ocr { 
		if v == value { return rune(A + byte(i)) }
	}
	return '?'
}

func (d Day8) rect(w int, h int) Day8 {
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			d.screen[y][x] = '#'
		}
	}
	return d
}

func (d Day8) rotateRow(row int, count int) Day8 {
	count = count % 50
	r := d.screen[row]

	left  := r[0:50-count]
	right := r[50-count:]
	
	newR := append(right, left...)
	
	for c := 0; c < 50; c++ { r[c] = newR[c] }

	d.screen[row] = r

	return d
}

func (d Day8) rotateCol(col int, count int) Day8 {
	count = count % 6
	
	r := [6]rune { d.screen[0][col], d.screen[1][col], d.screen[2][col], d.screen[3][col], d.screen[4][col], d.screen[5][col] }

	left  := r[0:6-count]
	right := r[6-count:]
	
	newR := append(right, left...)
	
	for c := 0; c < 6; c++ { 
		d.screen[c][col] = newR[c] 
	}

	return d
}

func (d Day8) execute(instruction string) Day8 {
	if strings.HasPrefix(instruction, "rect ") {
		wh := strings.Split(instruction[5:], "x")
		w, _ := strconv.Atoi(wh[0])
		h, _ := strconv.Atoi(wh[1])
		d = d.rect(w, h)
	} else if strings.HasPrefix(instruction, "rotate row y=") {
		rc := strings.Split(instruction[13:], " by ")
		row, _   := strconv.Atoi(rc[0])
		count, _ := strconv.Atoi(rc[1])
		d = d.rotateRow(row, count)
	} else if strings.HasPrefix(instruction, "rotate column x=") {
		cc := strings.Split(instruction[16:], " by ")
		col, _   := strconv.Atoi(cc[0])
		count, _ := strconv.Atoi(cc[1])
		d = d.rotateCol(col, count)
	}
	return d
}

func (d Day8) part1() int {
	for line := range readlines(8) {
		d = d.execute(line)
	}
	return d.count(-1)
}

func (d Day8) part2() string {
	for line := range readlines(8) {
		d = d.execute(line)
	}
	// d.print()

	var word strings.Builder

	for letter := 0; letter < 10; letter++ {
		word.WriteRune(d.getLetter(letter))
	}

	return word.String()
}

func (d Day8) run() {
	fmt.Println()
	fmt.Printf("--- Day 8 ---\n")
	fmt.Printf("Answer to day 8 part 1 is %v\n", d.part1())
	fmt.Printf("Answer to day 8 part 2 is %v\n", d.part2())
}
