package main

import (
	"fmt"
	"strings"

	"../aoc_utils"
)

func main() {
	result := 0
	input := aoc_utils.LoadInput("input")
	for _, i := range strings.Split(strings.Trim(input, "\n"), "\n") {
		if CheckIfValid(i) {
			result++
		}
	}
	fmt.Println("How many are valid yo? ", result)
}

func CheckIfValid(s string) bool {
	m := make(map[string]bool)
	for _, v := range strings.Fields(s) {
		if _, exist := m[v]; exist {
			return false
		} else {
			m[v] = true
		}
	}
	return true
}
