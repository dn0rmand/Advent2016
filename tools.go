package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readlines(day int) <-chan string {
	channel := make(chan string)

	go func() {
		file, err := os.Open(fmt.Sprintf("data/day%d.data", day))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			channel <- scanner.Text()
		}
		close(channel)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	return channel
}

func readItems(day int, separator string) <-chan string {
	channel := make(chan string)

	go func() {
		for line := range readlines(day) {
			for _,str := range strings.Split(line, separator) {
				channel <- strings.TrimSpace(str)
			}
		}
		close(channel)
	}()

	return channel
}
