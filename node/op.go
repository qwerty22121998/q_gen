package node

import (
	"q_gen/operator"
)

func AND(op ...Node) Node {
	return Operation{
		Type:     Logical,
		Operator: operator.AND,
		Operands: op,
	}
}

func OR(op ...Node) Node {
	return Operation{
		Type:     Logical,
		Operator: operator.OR,
		Operands: op,
	}
}

func NOT(op Node) Node {
	return Operation{
		Type:     Logical,
		Operator: operator.NOT,
		Operands: []Node{op},
	}
}

func CMP(op operator.Operator, value ...Node) Node {
	return Operation{
		Type:     Cmp,
		Operator: op,
		Operands: value,
	}
}

func FIELD(name string) Node {
	return Field(name)
}

func VAR(name string, data interface{}) Node {
	return &Variable{
		Name:  name,
		Value: data,
	}
}

func VAL(value interface{}) Node {
	return Value{
		Value: value,
	}
}
