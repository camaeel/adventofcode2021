package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	var elements []int

	var increase, prev int

	for l := 0; l < len(lines)-1; l += 1 {
		i, err := strconv.Atoi(lines[l])
		if err != nil {
			panic(err)
		}
		elements = append(elements, i)
	}

	for i := 0; i < len(elements)-2; i += 1 {
		curr := elements[i] + elements[i+1] + elements[i+2]
		if i > 0 {
			if curr > prev {
				increase += 1
			}
		}
		prev = curr
	}

	fmt.Printf("%d", increase)
}
