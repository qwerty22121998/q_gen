package node

import "q_gen/operator"

type Operation struct {
	Type     Type              `json:"type"`
	Operator operator.Operator `json:"operator"`
	Operands Group             `json:"operands"`
}

func (n Operation) String() string {

	val := make([]string, 0, len(n.Operands))
	for _, op := range n.Operands {
		val = append(val, op.String())
	}
	return n.Operator.Generate(val...)
}

func (Operation) T() Type {
	return Logical
}
