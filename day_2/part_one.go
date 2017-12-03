package main

import (
	"fmt"

	"../aoc_utils"
)

func main() {
	input := aoc_utils.ToMultiInt(aoc_utils.LoadInput("input"))
	// input := [][]int{{5, 1, 9, 5}, {7, 5, 3}, {2, 4, 6, 8}}
	fmt.Println("Checksum: ", GetChecksum(input))
}

func GetChecksum(input [][]int) (chk int) {
	for _, row := range input {
		min, max := row[0], row[0]
		for i := 1; i < len(row); i++ {
			if row[i] <= min {
				min = row[i]
			}
			if row[i] >= max {
				max = row[i]
			}
		}
		chk += max - min
	}
	return
}
