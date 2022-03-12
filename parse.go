package q_gen

import (
	"q_gen/node"
)

func Parse(raw string) *node.Tree {
	tree := &node.Tree{
		Head:   nil,
		Values: nil,
	}

	return tree

}
