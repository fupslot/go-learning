package main

import (
	"fmt"
	"time"
)

func trackTime(start time.Time) {
	diff := time.Since(start)
	fmt.Printf("%s", diff)
}

func main() {
	defer trackTime(time.Now())

	for i := 1e10; i > 0; i-- {
		// do something
	}
}
