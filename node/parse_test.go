package node

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree_Parse(t *testing.T) {
	data := `{"type":"logical","operator":"AND","operands":[{"type":"logical","operator":"OR","operands":[{"type":"comparison","operator":"=","operands":[{"type":"table","name":"gt.level"},1]},{"type":"comparison","operator":"=","operands":[{"type":"table","name":"gt.province"},"@04"]},{"type":"logical","operator":"NOT","operands":[{"type":"logical","operator":"AND","operands":[{"type":"table","name":"gt.level"},{"type":"table","name":"gt.level"},"gt.disable","gt.inactive"]}]}]}]}`
	tree, err := Parse(data)
	assert.Nil(t, err)
	fmt.Println(tree.String())
}
