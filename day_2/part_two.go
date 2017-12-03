package main

import (
	"fmt"

	"../aoc_utils"
)

func main() {
	input := aoc_utils.ToMultiInt(aoc_utils.LoadInput("input"))
	//input := [][]int{{5, 9, 2, 8}, {9, 4, 7, 3}, {3, 8, 6, 5}}
	fmt.Println("Checksum: ", GetChecksum(input))
}

func GetChecksum(input [][]int) (chk int) {
	for _, row := range input {
		for i := 0; i < len(row)-1; i++ {
			for j := i + 1; j < len(row); j++ {
				if row[i]%row[j] == 0 {
					chk += row[i] / row[j]
				} else if row[j]%row[i] == 0 {
					chk += row[j] / row[i]
				}
			}
		}
	}
	return
}
