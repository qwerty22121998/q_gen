package node

import (
	"q_gen/utils"
	"strings"
)

type Tree struct {
	Head   Node
	Values []*Variable
}

type ValueMap map[string]interface{}

func (v ValueMap) Get(path string) (interface{}, error) {
	paths := strings.Split(path, ".")
	if len(paths) == 0 {
		return v[path], nil
	}
	return utils.ExtractPath(v[path], paths[1:])
}

func (t *Tree) FillData(data ValueMap) error {
	for _, key := range t.Values {
		value, err := data.Get(key.Name)
		if err != nil {
			return err
		}
		key.Value = value
	}
	return nil
}
