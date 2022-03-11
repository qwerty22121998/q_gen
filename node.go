package q_gen

import (
	"fmt"
	"q_gen/operator"
)

type NodeType string

const (
	Logical    NodeType = "logical"
	Comparison NodeType = "comparison"
	Math       NodeType = "math"
	Variable   NodeType = "var"
	Value      NodeType = "value"
)

type Node interface {
	T() NodeType
	String() string
}

type NodeOP struct {
	Type     NodeType          `json:"type"`
	Operator operator.Operator `json:"operator"`
	Operands []Node            `json:"operands"`
}

func (n NodeOP) String() string {

	val := make([]string, 0, len(n.Operands))
	for _, op := range n.Operands {
		val = append(val, op.String())
	}
	return n.Operator.Generate(val...)
}

func (NodeOP) T() NodeType {
	return Logical
}

type NodeVariable string

func (NodeVariable) T() NodeType {
	return Variable
}

func (n NodeVariable) String() string {
	return string(n)
}

type NodeValue struct {
	Value interface{}
}

func (NodeValue) T() NodeType {
	return Value
}

func (n NodeValue) String() string {
	val, _ := n.MarshalJSON()
	return string(val)
}

func (n NodeValue) MarshalJSON() ([]byte, error) {
	switch n.Value.(type) {
	case string:
		return []byte(fmt.Sprintf(`"%v"`, n.Value)), nil
	default:
		return []byte(fmt.Sprint(n.Value)), nil
	}
}
