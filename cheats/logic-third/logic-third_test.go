package logicthird

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type flagtest struct {
	input  []int
	output int
}

func TestCountSumOfArray(t *testing.T) {
	flagtests := []flagtest{
		{
			input:  []int{1, 1, 2, 2},
			output: 3,
		},
		{
			input:  []int{1, 2, 1, 2},
			output: 0,
		},
		{
			input:  []int{1, 1, 1, 1},
			output: 4,
		},
		{
			input:  []int{3, 2, 2, 3},
			output: 5,
		},
	}
	for i, tt := range flagtests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := calculateSumOfArray(tt.input)
			assert.Equal(t, tt.output, result, "should be equal")
		})
	}
}
