package intfromto

func IntFromTo(from int, to int) []int {
	var arr []int
	i := from
	for i <= to {
		arr = append(arr, i)
		i++
	}
	return arr
}

func Int32FromTo(from int32, to int32) []int32 {
	var arr []int32
	i := from
	for i <= to {
		arr = append(arr, i)
		i++
	}
	return arr
}

func Int64FromTo(from int64, to int64) []int64 {
	var arr []int64
	i := from
	for i <= to {
		arr = append(arr, i)
		i++
	}
	return arr
}
