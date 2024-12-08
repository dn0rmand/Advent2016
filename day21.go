package main

import (
	"fmt"
	"strings"
)

// Day21 link https://adventofcode.com/2016/day/21
type Day21 struct {
}

func (d Day21) process(input []byte, command string, decrypt bool) []byte {
	l := len(input)

	switch {
		case strings.HasPrefix(command, "rotate right "): {
			value := int(command[13]-'0')
			if value > 0 {
				var right []byte
				var left []byte
				if decrypt {
					right = input[value:]
					left  = input[0:value]
				}	else {
					right = input[l-value:]
					left  = input[0:l-value]
				}

				input = []byte {}
				input = append(input, right...)
				input = append(input, left...)
			}
		}
		case strings.HasPrefix(command, "rotate left "): {
			value := int(command[12]-'0')
			if value > 0 {
				var right []byte
				var left  []byte

				if decrypt {
					right = input[l-value:]
					left  = input[0:l-value]
				} else {
					right = input[value:]
					left  = input[0:value]
				}

				input = []byte {}
				input = append(input, right...)
				input = append(input, left...)
			}
		}

		case strings.HasPrefix(command, "swap letter "): {
			letter1 := command[12]
			letter2 := command[len(command)-1]
			if letter1 != letter2 {
				for i := 0; i < l; i++ {
					switch input[i] {
						case letter1: input[i] = letter2
						case letter2: input[i] = letter1
					}
				}
			}
		}

		case strings.HasPrefix(command, "swap position "): {
			start := int(command[14]-'0')
			end   := int(command[len(command)-1]-'0')
			if start != end {
				input[start], input[end] = input[end], input[start]
			}
		}

		case strings.HasPrefix(command, "reverse positions "): {
			start := int(command[18]-'0')
			end   := int(command[len(command)-1]-'0')
			for i := 0; start+i < end-i; i++ {
				input[start+i] , input[end-i] = input[end-i], input[start+i]
			}
		}

		case strings.HasPrefix(command, "move position "): {
			from := int(command[14]-'0')
			to := int(command[len(command)-1]-'0')
			if decrypt { 
				from, to = to, from
			}

			c := input[from]
			if from > to {
				for i := from; i > to; i-- {
					input[i] = input[i-1]
				}
				input[to] = c
			} else if from < to {
				for i := from; i < to; i++ {
					input[i] = input[i+1]
				}
				input[to] = c
			} 
		}

		case strings.HasPrefix(command, "rotate based on position of letter "): {
			letter := command[len(command)-1]
			index  := strings.IndexByte(string(input), letter)			
			value  := index+1
			if value >= 5 { value++ }
			value %= l

			if decrypt {
					basedReverse := []int { 1, 1, 6, 2, 7, 3, 0, 4 }
					value = basedReverse[index]
			}

			var right []byte
			var left  []byte

			if (value > 0) {
				if decrypt {
					right = input[value:]
					left  = input[0:value]
				} else {
					right = input[l-value:]
					left  = input[0:l-value]
				}

				input = []byte {}
				input = append(input, right...)
				input = append(input, left...)
			}
		}
	}

	return input
}

func (d Day21) part1() string {
	password := []byte { 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h' }

	for line := range readlines(21) {
		password = d.process(password, line, false)
	}

	return string(password)
}

func (d Day21) part2() string {
	password := []byte { 'f', 'b', 'g', 'd', 'c', 'e', 'a', 'h' }
	commands := []string {}

	for line := range readlines(21) {
		commands = append(commands, line)
	}

	for i := len(commands); i > 0; i-- {
		password = d.process(password, commands[i-1], true)
	}

	return string(password)
}

func (d Day21) run() {
	fmt.Println()
	fmt.Printf("--- Day 21 ---\n")
	fmt.Printf("Answer to day 21 part 1 is %v\n", d.part1())
	fmt.Printf("Answer to day 21 part 2 is %v\n", d.part2())
}
