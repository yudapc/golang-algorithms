package interfaceimplementationsample

import "testing"

func TestImplementationOfInterface(t *testing.T) {
	if err := Implementation(); err != nil {
		t.Error("Error should be nil")
	}
}
