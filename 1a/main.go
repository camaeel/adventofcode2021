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
	var decrease, increase, prev int

	for l := range lines {
		if lines[l] != "" {
			i, err := strconv.Atoi(lines[l])
			if err != nil {
				panic(err)
			}
			if l > 0 {
				if i > prev {
					increase += 1
				} else if i < prev {
					decrease += 1
				}
			}
			prev = i
		}
	}

	fmt.Printf("%d", increase)
}
