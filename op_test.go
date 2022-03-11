package q_gen

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func Test_OP(t *testing.T) {
	data := AND(
		OR(
			CMP("=", FIELD("gt.level"), VAL(1)),
			CMP("=", FIELD("gt.province"), VAR("04", "hello")),
			NOT(AND(
				FIELD("gt.disable"),
				FIELD("gt.inactive"),
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
