package aoc_utils

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func LoadInput(path string) string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}

	return string(file)
}

func ToInt(s string) (out []int) {
	s = strings.Trim(s, "\n")
	for _, c := range s {
		v, err := strconv.Atoi(string(c))
		if err != nil {
			fmt.Println(err)
		} else {
			out = append(out, v)
		}
	}
	return
}

func ToMultiInt(s string) (out [][]int) {
	for _, i := range strings.Split(strings.Trim(s, "\n"), "\n") {
		row := make([]int, len(strings.Fields(i)))
		for k, j := range strings.Fields(i) {
			v, err := strconv.Atoi(j)
			if err != nil {
				fmt.Println(err)
			} else {
				row[k] = v
			}
		}
		out = append(out, row)
	}
	return
}
