package main

import (
	"fmt"
	"strconv"
	"strings"

	"../aoc_utils"
)

func main() {
	file := aoc_utils.LoadInput("input")
	m := make(map[string]int)

	for _, v := range strings.Split(strings.Trim(file, "\n"), "\n") {
		ins := strings.Fields(v)

		//Conditions and values
		amount, _ := strconv.Atoi(ins[2])
		cond := ins[4]
		comp := ins[5]
		conV, _ := strconv.Atoi(ins[6])

		//What i did here, needs to be erased from mankind history...   D: shit
		switch ins[1] {
		case "inc":
			switch comp {
			case ">":
				if m[cond] > conV {
					m[ins[0]] += amount
				}
			case "<":
				if m[cond] < conV {
					m[ins[0]] += amount
				}
			case ">=":
				if m[cond] >= conV {
					m[ins[0]] += amount
				}
			case "<=":
				if m[cond] <= conV {
					m[ins[0]] += amount
				}
			case "==":
				if m[cond] == conV {
					m[ins[0]] += amount
				}
			case "!=":
				if m[cond] != conV {
					m[ins[0]] += amount
				}
			}

		case "dec":
			switch comp {
			case ">":
				if m[cond] > conV {
					m[ins[0]] -= amount
				}
			case "<":
				if m[cond] < conV {
					m[ins[0]] -= amount
				}
			case ">=":
				if m[cond] >= conV {
					m[ins[0]] -= amount
				}
			case "<=":
				if m[cond] <= conV {
					m[ins[0]] -= amount
				}
			case "==":
				if m[cond] == conV {
					m[ins[0]] -= amount
				}
			case "!=":
				if m[cond] != conV {
					m[ins[0]] -= amount
				}
			}
		}
	}

	max := 0
	for k, v := range m {
		if v >= max {
			max = v
		}
		fmt.Printf("K: %s V: %d\n", k, v)
	}
	fmt.Println("\nMax: ", max)
}
