package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i) // Deferred functions are executed in LIFO order, so the above code prints: 4 3 2 1 0.
	}
}
