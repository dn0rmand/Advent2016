package main

import (
	"fmt"
	"sort"
	"strings"
)

type day11Object struct {
	id byte
	microship byte
	floor byte
}

type day11State struct {
	elevator byte
	objects []day11Object
}

func (d day11State) hasGenerator(o day11Object) bool {
	if o.microship == 0 { return true }
	for _, g := range d.objects {
		if (g.id == o.id && g.microship == 0) {
			return g.floor == o.floor
		}
	}
	return false
}

func (d day11State) dead(o day11Object) bool {
	if o.microship == 0 { return false }
	if d.hasGenerator(o) { return  false }
	for _, g := range d.objects {
		if (g.floor == o.floor && g.microship == 0 && g.id != o.id) {
			return true
		}
	}
	return false
}

func (d day11State) isValid() bool {
	for _, o := range d.objects {
		if d.dead(o) { return false }
	}
	return true
}

func (d day11State) success() bool {
	if d.elevator != 4 { return false }
	for _, o := range d.objects { 
		if o.floor != 4 { return false }
	}
	return true
}

func (d day11State) cloneOjects(i int, j int) []day11Object {
	if i > j {
		k := i
		i = j
		j = k
	}
	result := []day11Object {}

	if i > 0 {
		result = append(result, d.objects[0:i]...)
	}
	if j > i+1 {
		result = append(result, d.objects[i+1:j]...)
	}
	result = append(result, d.objects[j+1:]...)

	return result
}

func (d day11State) next() []day11State {
	var nextStates []day11State

	count := len(d.objects) 
	for i, o1 := range d.objects {
		if o1.floor != d.elevator { continue }

		up1   := day11Object { id: o1.id, microship: o1.microship, floor: o1.floor+1 }
		down1 := day11Object { id: o1.id, microship: o1.microship, floor: o1.floor-1 }
		
		// going up or down with only 1 object
		if d.elevator < 4 {
			next := day11State {
				elevator: d.elevator+1,
				objects: append(d.cloneOjects(i, i), up1),
			}
			if len(next.objects) != count  { println("NOT RIGHT 1") }

			if next.isValid() {
				nextStates = append(nextStates, next)
			}
		}
		if d.elevator > 1 {
			next := day11State {
				elevator: d.elevator-1,
				objects: append(d.cloneOjects(i, i), down1),
			}
			if len(next.objects) != count  { println("NOT RIGHT 2") }

			if next.isValid() {
				nextStates = append(nextStates, next)
			}
		}

		// going up  or  down with 2 objects
		for j, o2 := range d.objects {
			if o2.floor != d.elevator || j <= i { continue }

			if o2.id == o1.id || o2.microship == o1.microship { // need to be compatible
				up2   := day11Object { id: o2.id, microship: o2.microship, floor: o2.floor+1 }
				down2 := day11Object { id: o2.id, microship: o2.microship, floor: o2.floor-1 }

				if d.elevator < 4 {
					next := day11State {
						elevator: d.elevator+1,
						objects: append(d.cloneOjects(i, j), up1, up2),
					}
					if len(next.objects) != count  { println("NOT RIGHT 3") }

					if next.isValid() {
						nextStates = append(nextStates, next)
					}
				}
				if d.elevator > 1 {
					next := day11State {
						elevator: d.elevator-1,
						objects: append(d.cloneOjects(i, j), down1, down2),
					}
					if len(next.objects) != count  { println("NOT RIGHT 4") }

					if next.isValid() {
						nextStates = append(nextStates, next)
					}
				}
			}
		}
	}

	return nextStates
}

func (d day11State) getKey() string {
	sort.Slice(d.objects, func(i, j int) bool { 
		if d.objects[i].id == d.objects[j].id {
				return d.objects[i].microship < d.objects[j].microship 
		}
		return d.objects[i].id < d.objects[j].id 
	})

	values := []string { string([]byte { d.elevator + '0' }) }

	for i := 0; i < len(d.objects); i += 2 {
		o1 := d.objects[i]
		o2 := d.objects[i+1]

		var k []byte

		if o1.floor == o2.floor {
			k = []byte { 'P', '0'+o1.floor }
		} else {
			k = []byte { 'D', '0'+o1.floor, '0'+o2.floor }
		}

		values = append(values, string(k))
	}

	sort.Sort(sort.StringSlice(values))

	k := strings.Join(values, "")
	return k
}

// Day11 link https://adventofcode.com/2016/day/11
type Day11 struct {
}

func (d Day11) process(objects []day11Object) int {
	steps := 0
	visited := make(map[string]byte)

	states := []day11State { { objects: objects, elevator: 1 } }

	visited[states[0].getKey()] = 1
	for len(states) > 0 {
		steps++
		var nextStates []day11State

		for _, state := range states {
			for _, nextState := range state.next() {

				if nextState.success() { return steps }

				var k = nextState.getKey()
				if visited[k] == 1 { continue }
				visited[k] = 1
				nextStates = append(nextStates, nextState)
			}
		}

		states = nextStates
	}

	return -1
}

func (d Day11) loadInput() []day11Object {
	var objects []day11Object
	nextID := byte(0)
	names := make(map[string]byte)

	for line := range readlines(11) {		
		if strings.Index(line, "contains nothing relevant") > 0 { continue }

		line = line[:len(line)-1] // remove the .

		var floor byte 

		switch line[4:6] {
			case "fi": 
				floor = 1
			case "se": 
				floor = 2
			case "th": 
				floor = 3
			case "fo":
				floor = 4
		}

		line = line[strings.Index(line, "floor contains a ")+17:]

		line = strings.ReplaceAll(line, " a ", "")
		line = strings.ReplaceAll(line, ", and", ",")
		line = strings.ReplaceAll(line, " generator", "")
		line = strings.ReplaceAll(line, "-compatible microchip", "-M")
		entries := strings.Split(line, ",")

		for _, name := range entries {
			microship := byte(0)
			if strings.HasSuffix(name, "-M") { microship = 1 }
			if microship == 1 { name = name[:len(name)-2] }
			
			id := names[name]
			if id  == 0 {
				nextID++
				id = nextID
				names[name] = id
			}

			obj := day11Object { id: id, microship: microship, floor: floor }
			objects = append(objects, obj)
		}
	}

	return objects
}

func (d Day11) part1() int {
	objects := d.loadInput()
	return d.process(objects)
}

func (d Day11) part2() int {
	objects := d.loadInput()
	id := byte(len(objects) / 2)
	extra := []day11Object { 
		{ id: id+1, microship: 0, floor: 1},
		{ id: id+1, microship: 1, floor: 1},
		{ id: id+2, microship: 0, floor: 1},
		{ id: id+2, microship: 1, floor: 1},
	}
	objects = append(objects, extra...)
	return d.process(objects)
}

func (d Day11) run() {
	fmt.Println()
	fmt.Printf("--- Day 11 ---\n")
	fmt.Printf("Answer to day 11 part 1 is %v\n", d.part1())
	fmt.Printf("Answer to day 11 part 2 is %v\n", d.part2())
}
