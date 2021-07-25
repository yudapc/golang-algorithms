package logicthird

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type flagSecondTest struct {
	input  []int
	output map[string]int
}

func TestCountNumberTimes(t *testing.T) {
	flagtests := []flagSecondTest{
		{
			input: []int{1, 2, 3, 5, 1, 1},
			output: map[string]int{
				"1": 3,
				"2": 1,
				"3": 1,
				"5": 1,
			},
		},
		{
			input: []int{1, 2, 3, 2, 3, 5, 1, 1},
			output: map[string]int{
				"1": 3,
				"2": 2,
				"3": 2,
				"5": 1,
			},
		},
	}
	for i, tt := range flagtests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := countNumberTimes(tt.input)
			assert.Equal(t, tt.output, result, "should be equal")
		})
	}
}
