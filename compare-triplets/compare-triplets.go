package comparetriplets

func compareTriplets(a []int32, b []int32) []int32 {
	// Write your code here
	if len(a) != len(b) {
		return nil
	}
	alice := int32(0)
	bob := int32(0)
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			alice += 1
		}
		if a[i] < b[i] {
			bob += 1
		}
	}
	return []int32{alice, bob}
}
