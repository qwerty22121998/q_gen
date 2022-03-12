package node

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree_Parse(t *testing.T) {
	data := `{"type":"logical","operator":"AND","operands":[{"type":"logical","operator":"OR","operands":[{"type":"comparison","operator":"=","operands":["gt.level",1]},{"type":"comparison","operator":"=","operands":["gt.province","@04"]},{"type":"logical","operator":"NOT","operands":[{"type":"logical","operator":"AND","operands":["gt.disable","gt.inactive"]}]}]}]}`
	tree, err := Parse(data)
	assert.Nil(t, err)
	fmt.Println(tree.Head.String())

}
