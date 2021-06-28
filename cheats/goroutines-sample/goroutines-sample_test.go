package goroutinessample

import (
	"testing"
)

func TestGoroutines(t *testing.T) {
	ch := SampleGoroutines()
	t.Logf(<-ch, <-ch, <-ch, <-ch, <-ch)
}
