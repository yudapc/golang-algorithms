package logicthird

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestInteger(t *testing.T) {

	flagtests := []flagtest{
		{
			input:  []int{3, 5, 2, 4},
			output: 2,
		},
		{
			input:  []int{10, 4, 3, 2, 7, 1, 5, 8},
			output: 1,
		},
		{
			input:  []int{5, 4, 7, 8, 4, 5, 6, 3},
			output: 3,
		},
		{
			input:  []int{5, 6, 4, 8, 3, 9, 7, 8, 9, 2, 10, 11, 5, 5, 5, 3, 2, 3, 234, 2, 3, 5},
			output: 2,
		},
	}
	for i, tt := range flagtests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := smallestInteger(tt.input)
			assert.Equal(t, tt.output, result, "should be equal")
		})
	}
}
