package main

import (
	"fmt"

	"../aoc_utils"
)

func main() {
	file := aoc_utils.LoadInput("input")
	input := aoc_utils.ToInt(file)
	sum := find_sum(input)
	fmt.Println("Captcha: ", sum)
}

func find_sum(nums []int) (sum int) {
	jump := len(nums) / 2
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if jump+i < len(nums) {
			if nums[i] == nums[i+jump] {
				m[nums[i]] += 2
			}
		}
	}

	for k, v := range m {
		sum += k * v
	}

	return
}
