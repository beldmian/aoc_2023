package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	ss := strings.Split(string(f), "\n")
	sum := 0
	for _, game := range ss {
		sets := strings.Split(strings.Split(game, ": ")[1], "; ")
		redMin := 0
		greenMin := 0
		blueMin := 0
		for _, set := range sets {
			gems := strings.Split(set, ", ")
			for _, gem := range gems {
				gem_data := strings.Split(gem, " ")
				gem_number, _ := strconv.Atoi(gem_data[0])
				gem_color := gem_data[1]
				if gem_color == "red" && gem_number > redMin {
					redMin = gem_number
				} else if gem_color == "green" && gem_number > greenMin {
					greenMin = gem_number
				} else if gem_color == "blue" && gem_number > blueMin {
					blueMin = gem_number
				}
			}
		}
		sum += redMin * greenMin * blueMin
	}
	println(sum)
}
