package main

import (
	"fmt"
	"sort"
	"strings"

	"../aoc_utils"
)

func main() {

	result := 0
	input := aoc_utils.LoadInput("input")
	for _, s := range strings.Split(strings.Trim(input, "\n"), "\n") {
		if CheckIfValid(s) {
			result++
		}
	}

	fmt.Println("How many are valid YO? ", result)
}

func CheckIfValid(s string) bool {
	m := make(map[string]bool)
	for _, v := range strings.Fields(s) {
		w := strings.Split(v, "")
		sort.Strings(w)
		sorted := strings.Join(w, "")

		if _, exist := m[sorted]; exist {
			return false
		}
		m[sorted] = true
	}
	return true
}
