package makesample

import "testing"

func TestCreationSliceOfInt(t *testing.T) {
	result := CreationSliceOfInt()
	if len(result) < 1 {
		t.Error("Error creation slice of int", len(result))
	}
}
