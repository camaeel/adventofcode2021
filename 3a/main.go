package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	numOfLines := len(lines) - 1
	var ones []int

	for l := 0; l < len(lines)-1; l += 1 {
		if l == 0 {
			ones = make([]int, len(lines[l]))
		}
		fields := strings.Split(lines[l], "")
		for f := range fields {
			if fields[f] == "1" {
				ones[f] += 1
			}
		}
	}

	var gammaStr, epsilonStr string

	for f := range ones {
		if ones[f] > numOfLines/2 {
			gammaStr += "1"
			epsilonStr += "0"
		} else {
			gammaStr += "0"
			epsilonStr += "1"
		}
	}

	gamma := bin2int(gammaStr)
	epsilon := bin2int(epsilonStr)

	fmt.Printf("%d", gamma*epsilon)
}

func bin2int(str string) int {
	fields := strings.Split(str, "")
	ret := 0
	for i := range fields {
		ret *= 2
		if fields[i] == "1" {
			ret += 1
		}
	}
	return ret
}
