package main

import (
	"crypto/md5"
	"fmt"
	"strings"
	"strconv"
)

// Day5 link https://adventofcode.com/2016/day/5
type Day5 struct {
}

func (d Day5) getHashes() chan []byte {
	channel := make(chan []byte)
	prefix := readline(5)
	hasher := md5.New()

	go func() {
		for i := 0; ; i++ {
			hasher.Reset()
			hasher.Write([]byte(prefix))
			hasher.Write([]byte(strconv.Itoa(i)))
			var hash = hasher.Sum(nil)
			if hash[0] == 0 && hash[1] == 0 && hash[2] <= 0xF {
				channel <- hash
			}
		} 
	}()

	return channel
}

func (d Day5) runesToString(runes [8]rune) string {
	var result strings.Builder

	for _,r := range runes { result.WriteRune(r) }

	return result.String()
}

func (d Day5) parts() (string, string) {
	chars := [16] rune { '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f' }

	var password1 strings.Builder
	var password2 = [8]rune { '*', '*', '*', '*', '*', '*', '*', '*' }
	var found     = 0

	channel := d.getHashes()

	for b := range channel {	
		changed := false	
		b2 := b[2]
		b3 := (b[3] >> 4) & 0x0F

		if password1.Len() < 8 {
			password1.WriteRune(chars[b2])
		}

		if b2 < 8 && password2[b2] == '*' {
			password2[b2] = chars[b3]
			found++
			changed = true
		}

		if changed { 
			fmt.Print("\r", d.runesToString(password2))
		}

		if password1.Len()==8 && found == 8 { break }
	}

	close(channel)

	return password1.String(), d.runesToString(password2)
}

func (d Day5) run() {
	fmt.Println()
	fmt.Printf("--- Day 5 ---\n")
	var pwd1, pwd2 = d.parts()
	fmt.Printf("Answer to day 5 part 1 is %v\n", pwd1)
	fmt.Printf("Answer to day 5 part 2 is %v\n", pwd2)
}
