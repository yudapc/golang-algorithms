package goroutinessample

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutines(t *testing.T) {
	ch := SampleGoroutines()
	t.Logf(<-ch, <-ch, <-ch, <-ch, <-ch)
}

// Sample 2 Test
func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")
	fmt.Println("Ups 2")
	go RunHelloWorld()
	fmt.Println("Ups 3")
	time.Sleep(1 * time.Second)
}
