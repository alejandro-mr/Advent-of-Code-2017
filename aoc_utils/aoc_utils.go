package aoc_utils

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"../data_structures/tree"
)

func LoadInput(path string) string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}

	return string(file)
}

func ParseTreeInput(s string) (out []tree.Node) {
	b := strings.Split(strings.Trim(s, "\n"), "\n")
	for _, v := range b {
		reNode := regexp.MustCompile(`[a-zA-Z]+ \([0-9]+\)`)
		nData := reNode.FindAllString(v, -1)[0]

		reChildren := regexp.MustCompile(`-> ([a-z, ]+)+`)
		cData := reChildren.FindStringSubmatch(v)
		if len(cData) > 0 {
			cData = strings.Fields(strings.Replace(cData[1], ",", "", -1))
		}

		reNum := regexp.MustCompile(`\D`)
		weight, _ := strconv.Atoi(reNum.ReplaceAllString(nData, ""))

		reNam := regexp.MustCompilePOSIX(`[^a-z]`)
		name := reNam.ReplaceAllString(nData, "")

		node := tree.Node{
			Parent: nil,
			Value: tree.Entry{
				Name:   name,
				Weight: weight,
			},
			Children: nil,
		}
		children := make([]*tree.Node, 0)
		for _, v := range cData {
			children = append(children, &tree.Node{
				Parent: &node,
				Value: tree.Entry{
					Name:   v,
					Weight: 0,
				},
				Children: nil,
			})
		}
		node.Children = children
		out = append(out, node)
	}
	return
}

func ToSliceInt(s string) (out []int) {
	s = strings.Trim(s, "\n")
	for _, c := range strings.Fields(s) {
		v, err := strconv.Atoi(string(c))
		if err != nil {
			fmt.Println(err)
		} else {
			out = append(out, v)
		}
	}
	return
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
