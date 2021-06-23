package arrayleftrotation

func ArrayLeftRotation(a []int32, d int32) []int32 {
	rotation := d % int32(len(a))
	if rotation == 0 {
		return a
	}

	arr1 := make([]int32, 0)
	arr2 := make([]int32, 0)

	for i := range a {
		if i < int(rotation) {
			arr1 = append(arr1, a[i])
		} else {
			arr2 = append(arr2, a[i])
		}
	}

	return append(arr2, arr1...)
}

func ArrayLeftRotationSimple(a []int32, d int32) []int32 {
	totalLen := len(a)
	first := a[:d]
	last := a[d:totalLen]
	return append(last, first...)
}
