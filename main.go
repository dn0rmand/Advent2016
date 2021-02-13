package main

import (
	"fmt"
	"time" 
)

type day interface { 
	run() 
}

func dumpTiming(msg string, start time.Time, end time.Time) {
	duration := end.Sub(start)
	secs := int(duration.Seconds())
	ms   := int(duration.Milliseconds()) % 1000
	us   := int(duration.Microseconds()) % 1000

	fmt.Printf("%vExecuted in %v sec, %v ms and %v μs\n", msg, secs, ms, us)
}

func exec(d day) {
	start := time.Now()
	d.run()
	end := time.Now()

	dumpTiming("", start, end)
}

func main() {
	println("############################")
	println("# Avent of Code 2016 in Go #")
	println("############################")

	start := time.Now()

	exec(Day1{})
	exec(Day2{})
	exec(Day3{})
	exec(Day4{})
	exec(Day5{}) // very slow due to MD5
	exec(Day6{})
	exec(Day7{})
	exec(Day8{})
	exec(Day9{})
	exec(Day10{})
	exec(Day11{})
	exec(Day12{})
	exec(Day13{})
	exec(Day14{}) // slow due to MD5
	exec(Day15{})
	exec(Day16{})
	exec(Day17{})
	exec(Day18{})
	exec(Day19{})
	exec(Day20{})
	exec(Day21{})
	exec(Day22{})
	exec(Day23{})
	exec(Day24{})
	exec(Day25{})

	end := time.Now()

	dumpTiming("\nAll 25 days ", start, end)
}
