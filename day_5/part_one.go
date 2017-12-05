package main

import (
	"fmt"

	"../aoc_utils"
)

func main() {
	//input := []int{0, 3, 0, 1, -3}
	file := aoc_utils.LoadInput("input")
	input := aoc_utils.ToSliceInt(file)

	fmt.Println("Steps taken: ", CalculateSteps(input))
}

func CalculateSteps(ins []int) (steps int) {
	for i := 0; i < len(ins); i++ {
		if ins[i] != 0 {
			c := ins[i]
			ins[i]++
			i += c - 1
		} else {
			ins[i]++
			i--
		}
		steps++
	}
	return
}
