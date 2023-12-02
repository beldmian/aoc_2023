package main

import (
	"os"
	"strings"
)

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	ss := strings.Split(string(f), "\n")
	sum := 0
	for _, s := range ss {
		firstDigit := -1
		lastDigit := -1
		for len(s) > 0 && firstDigit == -1 {
			if strings.HasPrefix(s, "zero") || strings.HasPrefix(s, "0") {
				firstDigit = 0
			} else if strings.HasPrefix(s, "one") || strings.HasPrefix(s, "1") {
				firstDigit = 1
			} else if strings.HasPrefix(s, "two") || strings.HasPrefix(s, "2") {
				firstDigit = 2
			} else if strings.HasPrefix(s, "three") || strings.HasPrefix(s, "3") {
				firstDigit = 3
			} else if strings.HasPrefix(s, "four") || strings.HasPrefix(s, "4") {
				firstDigit = 4
			} else if strings.HasPrefix(s, "five") || strings.HasPrefix(s, "5") {
				firstDigit = 5
			} else if strings.HasPrefix(s, "six") || strings.HasPrefix(s, "6") {
				firstDigit = 6
			} else if strings.HasPrefix(s, "seven") || strings.HasPrefix(s, "7") {
				firstDigit = 7
			} else if strings.HasPrefix(s, "eight") || strings.HasPrefix(s, "8") {
				firstDigit = 8
			} else if strings.HasPrefix(s, "nine") || strings.HasPrefix(s, "9") {
				firstDigit = 9
			} else {
				s = s[1:]
			}
		}
		for len(s) > 0 && lastDigit == -1 {
			if strings.HasSuffix(s, "zero") || strings.HasSuffix(s, "0") {
				lastDigit = 0
			} else if strings.HasSuffix(s, "one") || strings.HasSuffix(s, "1") {
				lastDigit = 1
			} else if strings.HasSuffix(s, "two") || strings.HasSuffix(s, "2") {
				lastDigit = 2
			} else if strings.HasSuffix(s, "three") || strings.HasSuffix(s, "3") {
				lastDigit = 3
			} else if strings.HasSuffix(s, "four") || strings.HasSuffix(s, "4") {
				lastDigit = 4
			} else if strings.HasSuffix(s, "five") || strings.HasSuffix(s, "5") {
				lastDigit = 5
			} else if strings.HasSuffix(s, "six") || strings.HasSuffix(s, "6") {
				lastDigit = 6
			} else if strings.HasSuffix(s, "seven") || strings.HasSuffix(s, "7") {
				lastDigit = 7
			} else if strings.HasSuffix(s, "eight") || strings.HasSuffix(s, "8") {
				lastDigit = 8
			} else if strings.HasSuffix(s, "nine") || strings.HasSuffix(s, "9") {
				lastDigit = 9
			} else {
				s = s[:len(s)-1]
			}
		}
		println(firstDigit, lastDigit)
		sum += 10*firstDigit + lastDigit
	}

	println(sum)
}
