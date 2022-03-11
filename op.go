package q_gen

import (
	"q_gen/node"
	"q_gen/operator"
)

func AND(op ...node.Node) node.Node {
	return node.Operation{
		Type:     node.Logical,
		Operator: operator.AND,
		Operands: op,
	}
}

func OR(op ...node.Node) node.Node {
	return node.Operation{
		Type:     node.Logical,
		Operator: operator.OR,
		Operands: op,
	}
}

func NOT(op node.Node) node.Node {
	return node.Operation{
		Type:     node.Logical,
		Operator: operator.NOT,
		Operands: []node.Node{op},
	}
}

func CMP(op operator.Operator, value ...node.Node) node.Node {
	return node.Operation{
		Type:     node.Cmp,
		Operator: op,
		Operands: value,
	}
}

func FIELD(name string) node.Node {
	return node.Field(name)
}

func VAR(name string, data interface{}) node.Node {
	return &node.Variable{
		Name: name,
		Value: data,
	}
}

func VAL(value interface{}) node.Node {
	return node.Value{
		Value: value,
	}
}
