package tree

import "fmt"

type Tree struct {
	root  *Node
	depth int
}

func (t *Tree) Insert(parent *Node, n *Node, s *Node) {
	if t.root == nil {
		t.root = n
		t.depth += 1
		//fmt.Println("Inserting root")
		return
	} else if parent != nil {
		//fmt.Println("Inserting child in parent node")
		if parent == t.root {
			t.root.Children = append(t.root.Children, n)
			//fmt.Println("Found that it belongs to root")
			return
		}

		if s == nil {
			s = t.root
		}
		for _, sn := range s.Children {
			//fmt.Printf("Looking for parent on: %v p: %p\n", sn, sn)
			if sn.equals(*parent) {
				//fmt.Println("Doing recursive insertion")

				//Just removed this section since my code already does that
				sn.Children = append(sn.Children, n)
				return
			}
			//fmt.Println("Not found")
			t.Insert(parent, n, sn)
		}
	} else {
		t.root.Children = append(t.root.Children, n)
		t.depth += 1
		//fmt.Println("Inserting child of root")
		return
	}
}

func (t Tree) GetRoot() Node {
	return *t.root
}

func (t Tree) PrintTree() string {
	out := fmt.Sprint("\t-- Printing TREE --\n\n")
	if t.root != nil {
		out += fmt.Sprintf("%v\n", t.root.Value.Name)
		out += fmt.Sprint("\\\n\\\n")
		for _, c := range t.root.Children {
			out += fmt.Sprint(c.PrintNode(1))
		}
		return out
	}
	return "tree is empty"
}
