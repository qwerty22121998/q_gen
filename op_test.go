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
			CMP("=", VAR("gt.level"), VAL(1)),
			CMP("=", VAR("gt.province"), VAL("04")),
		),
	)

	j, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data)
	fmt.Println(string(j))
}
