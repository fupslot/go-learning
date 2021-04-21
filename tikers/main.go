package main

import (
	"fmt"
	"time"
)

func main() {
	tiker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-tiker.C:
				fmt.Println("Tick at", t.Format(time.RFC3339))
			}
		}
	}()

	time.Sleep(5 * time.Second)
	tiker.Stop()
	done <- true
	fmt.Println("Tiker stopped")
}
