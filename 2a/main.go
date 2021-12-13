package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	var depth, horizontal int

	for l := 0; l < len(lines)-1; l += 1 {
		fields := strings.Split(lines[l], " ")

		value, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}
		if fields[0] == "up" {
			depth -= value
		} else if fields[0] == "down" {
			depth += value
		} else if fields[0] == "forward" {
			horizontal += value
		}
	}

	fmt.Printf("%d", depth*horizontal)
}
