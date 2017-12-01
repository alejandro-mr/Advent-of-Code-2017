package aoc_utils

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ToInt(s string) (out []int) {
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

func LoadInput(path string) string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}

	return strings.TrimSuffix(string(file), "\n")
}
