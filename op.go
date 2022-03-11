package q_gen

import "q_gen/operator"

func AND(op ...Node) Node {
	return NodeOP{
		Type:     Logical,
		Operator: operator.AND,
		Operands: op,
	}
}

func OR(op ...Node) Node {
	return NodeOP{
		Type:     Logical,
		Operator: operator.OR,
		Operands: op,
	}
}

func NOT(op Node) Node {
	return NodeOP{
		Type:     Logical,
		Operator: operator.NOT,
		Operands: []Node{op},
	}
}

func CMP(op operator.Operator, value ...Node) Node {
	return NodeOP{
		Type:     Comparison,
		Operator: op,
		Operands: value,
	}
}

func FIELD(name string) Node {
	return NodeField(name)
}

func VAR(name string) Node {
	return NodeVariable(name)
}

func VAL(value interface{}) Node {
	return NodeValue{
		Value: value,
	}
}
