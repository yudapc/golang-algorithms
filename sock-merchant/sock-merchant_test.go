package sockmerchant

import "testing"

func TestSockMerchant(t *testing.T) {
	input := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	expected := int32(3)
	result := sockMerchant(int32(len(input)), input)
	if result != expected {
		t.Error("The result should be ", expected)
		t.Error("Received ", result)
	}
}
