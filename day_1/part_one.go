package main

import (
	"fmt"

	"../aoc_utils"
)

func main() {
	//input := []int{1, 1, 1, 1}
	//input := []int{1, 1, 2, 2}
	//input := []int{1, 2, 3, 4}
	//input := []int{9, 1, 2, 1, 2, 1, 2, 9}

	file := aoc_utils.LoadInput("input")
	input := aoc_utils.ToInt(file)

	sum := find_sum(input)
	fmt.Println("Captcha: ", sum)
}

func find_sum(nums []int) (sum int) {
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			sum += nums[i]
		}
	}

	if nums[0] == nums[len(nums)-1] {
		sum += nums[0]
	}

	return
}
