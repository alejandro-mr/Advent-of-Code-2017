package tree

import "fmt"

type Node struct {
	Parent   *Node
	Children []*Node
	Value    Entry
}

type Entry struct {
	Name   string
	Weight int
	Total  int
}

func (e Entry) empty() bool {
	if e.Name == "" && e.Weight == 0 {
		return true
	}
	return false
}

func (n Node) empty() bool {
	if n.Parent == nil && len(n.Children) < 1 && n.Value.empty() {
		return true
	}
	return false
}

func (n Node) equals(c Node) bool {
	if n.Parent == c.Parent && len(n.Children) == len(c.Children) && n.Value.Name == c.Value.Name && n.Value.Weight == n.Value.Weight {
		return true
	}
	return false
}

func (n Node) PrintNode(l int) string {
	padding := ""
	for i := l; i > 0; i-- {
		padding += "\t"
	}
	out := fmt.Sprintf("---------> %v (%v)\n", n.Value.Name, n.Value.Total)
	if len(n.Children) > 0 {
		out += padding + fmt.Sprint("\\\n")
		out += padding + fmt.Sprint("\\\n")
		l++
		for _, sn := range n.Children {
			out += padding + sn.PrintNode(l)
		}
	}
	return out
}
