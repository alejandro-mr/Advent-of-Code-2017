package main

import (
	"fmt"

	"../aoc_utils"
	"../data_structures/tree"
)

func main() {
	file := aoc_utils.LoadInput("input")
	input := aoc_utils.ParseTreeInput(file)

	//I used this code to visually validate if input is correctly parsed.
	/*
		o := ""

		for _, v := range input {
				o += fmt.Sprintf("%v (%v)", v.Value.Name, v.Value.Weight)
				o += fmt.Sprint("\tChildren: ")
				for _, c := range v.Children {
					o += fmt.Sprintf("%v,", c.Value.Name)
				}
				o += "\n"
		}
		fmt.Print(o)
	*/
	t := generateTree(input)
	fmt.Println(t.GetRoot().Value.Name)
}

//This method gets a list of nodes, and returns the tree structure
func generateTree(nodes []tree.Node) (t tree.Tree) {
	m := make(map[string]tree.Node)
	for _, n := range nodes {
		m[n.Value.Name] = n
	}

	root := findRoot(nodes)

	t.Insert(nil, generateBranches(&root, m), nil)
	return
}

func findRoot(nodes []tree.Node) tree.Node {
	var root tree.Node
	cm := make(map[string]bool)

	for _, n := range nodes {
		for _, c := range n.Children {
			cm[c.Value.Name] = true
		}
	}

	for _, n := range nodes {
		if _, exist := cm[n.Value.Name]; !exist && len(n.Children) > 0 {
			root = n
		}
	}

	return root
}

func generateBranches(n *tree.Node, m map[string]tree.Node) *tree.Node {
	if len(n.Children) > 0 {
		for k, c := range n.Children {
			if f, exist := m[c.Value.Name]; exist {
				n.Children[k] = &f
				n.Children[k].Parent = n
			}
			generateBranches(n.Children[k], m)
		}
	}

	return n
}
