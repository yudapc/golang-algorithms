package makesample

import "fmt"

func CreationSliceOfInt() []int {
	n := make([]int, 3)
	n[0] = 1
	n[2] = 2
	n[3] = 3
	fmt.Println(n)

	return n
}
