package main

import (
	"fmt"

	"../aoc_utils"
)

func main() {
	file := aoc_utils.LoadInput("input")

	groups, score, level := 0, 0, 0
	trash_started := false
	cancelling := false

	for _, v := range file {
		c := string(v)

		if !cancelling {
			//We are inside garbage area
			if trash_started {

				if c == "!" {
					cancelling = true
				}
				if c == ">" {
					trash_started = false
				}
			} else {
				if c == "{" {
					level++
					score += 1 * level
				}

				if c == "<" {
					trash_started = true
				}

				if c == "}" {
					groups++
					level--
				}
			}

		} else {
			cancelling = false
		}
	}

	fmt.Println("Groups: ", groups)
	fmt.Println("Score: ", score)
}
