package main

import (
	"fmt"
	"strconv"
	"strings"	
)

type day10Value struct {
	value int
	next *day10Value
}

type day10List struct {
	front *day10Value
	count int	
}

func (d *day10List) push(value int) {
	d.front = &day10Value{ value: value, next: d.front }
	d.count++
}

func (d *day10List) pop() int {
	if d.front != nil {
		v := d.front.value
		d.front = d.front.next
		d.count--
		return v
	} else {
		return -1
	}
}

type day10Bot struct {
	number int
	low int
	high int
	values *day10List
}

// Day10 link https://adventofcode.com/2016/day/10
type Day10 struct {
	bots map[int]*day10Bot
	outputs map[int]*day10List
}

func (d Day10) loadData() Day10 {
	distribution := make(map[int]int)
	bots := make(map[int]*day10Bot)
	outputs := make(map[int]*day10List)

	for line := range readlines(10) {
		if strings.HasPrefix(line, "value") {
			line = strings.Replace(line, "value ", "", 1)
			line = strings.Replace(line, " goes to bot ", ",", 1);
			info := strings.Split(line, ",")
			v, _ := strconv.Atoi(info[0])
			b, _ := strconv.Atoi(info[1])
			distribution[v] = b
		} else if strings.HasPrefix(line, "bot ") {
			line = strings.Replace(line, "bot ", "", 1)
			line = strings.Replace(line, " gives low to bot ", ",",1)
			line = strings.Replace(line, " gives low to output ", ",-1",1)
			line = strings.Replace(line, " and high to bot ", ",",1)
			line = strings.Replace(line, " and high to output ", ",-1",1)
			info:= strings.Split(line, ",")
			b,_ := strconv.Atoi(info[0])
			l,_ := strconv.Atoi(info[1])
			h,_ := strconv.Atoi(info[2])

			bot := day10Bot { number: b, low: l, high: h, values: &day10List {} } 
			if (bot.low < 0) { outputs[-bot.low] = &day10List {} }
			if (bot.high < 0) { outputs[-bot.high] = &day10List {} }

			bots[b] = &bot
		}
	}

	for value, bot := range distribution {
		bots[bot].values.push(value)
	}

	d.bots = bots
	d.outputs = outputs
	return d
}

func (d Day10) process() (int, int) {
	done := false

	answer1 := -1

	for ! done {
		done = true
		for _, bot := range d.bots {
			if bot.values.count >= 2 {
				v1, v2 := bot.values.pop() , bot.values.pop()

				if v1 > v2 { 
					v:= v1
					v1 = v2
					v2 = v 
				}

				if v1 == 17 && v2 == 61 { 
					if answer1 >= 0 && answer1 != bot.number { println("More than 1 bot ")}
					answer1 = bot.number
				}

				if bot.low >= 0 {
					d.bots[bot.low].values.push(v1)
					done = false
				} else {
					d.outputs[-bot.low].push(v1)
				}
				
				if bot.high >= 0 {
					d.bots[bot.high].values.push(v2)
					done = false
				} else {
					d.outputs[-bot.high].push(v2)
				}
			}
		}
	}

	answer2 := d.outputs[10].pop() * d.outputs[11].pop() * d.outputs[12].pop()
	return answer1, answer2
}

func (d Day10) parts() (int, int) {
	d = d.loadData()
	return d.process()
}

func (d Day10) run() {
	fmt.Println()
	fmt.Printf("--- Day 10 ---\n")
	part1, part2 := d.parts()
	fmt.Printf("Answer to day 10 part 1 is %v\n", part1)
	fmt.Printf("Answer to day 10 part 2 is %v\n", part2)
}
