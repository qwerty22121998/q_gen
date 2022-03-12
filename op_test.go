package q_gen

import (
	"encoding/json"
	"fmt"
	"log"
	"q_gen/node"
	"testing"
)

func Test_OP(t *testing.T) {
	data := node.AND(
		node.OR(
			node.CMP("=", node.FIELD("gt.level"), node.VAL(1)),
			node.CMP("=", node.FIELD("gt.province"), node.VAR("04", "hello")),
			node.NOT(node.AND(
				node.FIELD("gt.disable"),
				node.FIELD("gt.inactive"),
			),
			),
		),
	)

	j, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data)
	fmt.Println(string(j))
}
