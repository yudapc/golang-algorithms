package goroutinessample

import "fmt"

func SampleGoroutines() chan string {
	// Channel
	ch := make(chan string)

	// Start concurrent routines
	go push("Joko", ch)
	go push("Budi", ch)
	go push("Eko", ch)
	go push("Joni", ch)
	go push("Asep", ch)

	// Read 5 results
	// (Since our goroutines are concurrent,
	// the order isn't guaranteed!)
	// fmt.Println(<-ch, <-ch, <-ch, <-ch, <-ch)
	return ch
}

func push(name string, ch chan string) {
	msg := "Hey " + name
	ch <- msg
}

/// Sample 2
func RunHelloWorld() {
	fmt.Println("Hello World")
}
