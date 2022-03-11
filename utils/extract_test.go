package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtract(t *testing.T) {
	type example = struct {
		A int
		B string
		C []int
		D struct {
			E int
		}
	}

	a := example{
		A: 1,
		B: "abc",
		C: []int{1, 2, 3},
		D: struct{ E int }{E: 100},
	}

	field, err := Extract(a, "D.E")

	assert.Nil(t, err)
	assert.Equal(t, a.D.E, field)
	field, err = Extract(a, "D.F")
	assert.Error(t, err)
}
