package main

import (
	"crypto/md5"
	"fmt"
	"strconv"	
)

type day14Hash struct {
	data []byte
}

// Day14 link https://adventofcode.com/2016/day/14
type Day14 struct {
	buffer []day14Hash
	channel chan day14Hash
	part1Done bool
	part2Done bool
}

func (d *Day14) part1HashChannel() chan day14Hash {
	channel := make(chan day14Hash, 10)	
	prefix 	:= []byte(readline(14))
	hasher 	:= md5.New()

	go func() {
		for i := 0; ; i++ {
			if d.part1Done { 
				break 
			}
			hasher.Reset()
			hasher.Write(prefix)
			hasher.Write([]byte(strconv.Itoa(i)))
			var hash = hasher.Sum(nil)
			channel <- day14Hash { data: hash }
		} 

		channel <- day14Hash{}
		close(channel)
	}()

	return channel
}

func (d *Day14) part2HashChannel() chan day14Hash {
	channel := make(chan day14Hash, 10)	
	prefix 	:= []byte(readline(14))
	hasher 	:= md5.New()

	hex := "0123456789abcdef"

	go func() {
		for i := 0; ; i++ {
			if d.part2Done { break }

			hasher.Reset()
			hasher.Write(prefix)
			hasher.Write([]byte(strconv.Itoa(i)))
			var hash = hasher.Sum(nil)
			for i := 0; i < 2016; i++ {
				hasher.Reset()
				for _, b := range hash {
					values := []byte { hex[(b & 0xF0) >> 4], hex[b & 0x0F] }
					hasher.Write(values)
				}
				hash = hasher.Sum(nil)
			}
			channel <- day14Hash { data: hash }
		} 
		
		channel <- day14Hash{}
		close(channel)
	}()

	return channel
}

func (d Day14) hasConsecutive(hash []byte, previous byte) bool {
	c := 0
	for _, b := range hash {
		b1 := (b & 0xF0) >> 4
		if b1 == previous {
			c++
			if c == 5 { 
				return true 
			}
		} else {
			c = 0
		}

		b2 := b & 0x0F
		if b2 == previous {
			c++
			if c == 5 { 
				return true 
			}
		} else {
			c = 0
		}
	}

	return false
}

func (d *Day14) peek(index int) day14Hash {
	for index >= len(d.buffer) {
		value := <- d.channel
		d.buffer = append(d.buffer, value)
	}
	return d.buffer[index]
}

func (d *Day14) pop() day14Hash {
	if len(d.buffer) == 0 {
		value := <- d.channel
		return value
	}

	value := d.buffer[0];
	d.buffer = d.buffer[1:]
	return value
}


func (d *Day14) isValid(index int, value byte) bool {
	for count := 0; count < 1000; count++ {
		current := d.peek(count);

		if d.hasConsecutive(current.data, value) { 
			return true 
		}
	}

	return false
}

func (d *Day14) isKey(index int) bool {
	hash := d.pop()

	previous := (hash.data[0] & 0xF0) >> 4
	count    := 0

	for _, b := range hash.data {
		b1 := (b & 0xF0) >> 4
		if b1 == previous { 
			count++
			if count == 3 { 
				return d.isValid(index, b1)
			}
		} else {
			previous = b1 
			count = 1
		}

		b1 = b & 0xF
		if b1 == previous { 
			count++
			if count == 3 {
				return d.isValid(index, b1)
			}
		} else {
			previous = b1 
			count = 1
		}
	}

	return false
}

func (d *Day14) getNextIndex(start int) int {
	for index := start; ; index++ {
		if d.isKey(index) { return index }
	}
}

func (d Day14) part1() int {
	d.channel = d.part1HashChannel()

	index := -1
	for i := 1; i <= 64; i++ {
		index = d.getNextIndex(index+1)
	}

	d.part1Done = true

	// purge
	for value := range d.channel {
		if value.data == nil { break }
	}
	return index
}

func (d Day14) part2() int {
	d.channel = d.part2HashChannel()
	// d.buffer = []day14Hash {}

	index := -1
	for i := 1; i <= 64; i++ {
		index = d.getNextIndex(index+1)
	}

	d.part2Done = true
	// purge
	for value := range d.channel {
		if value.data == nil { break }
	}
	return index
}

func (d Day14) run() {
	fmt.Println()
	fmt.Printf("--- Day 14 ---\n")

	fmt.Printf("Answer to day 14 part 1 is %v\n", d.part1())
	fmt.Printf("Answer to day 14 part 2 is %v\n", d.part2())
}
