package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 1)
	done := make(chan bool)

	go func() {
		for { // run indefinetely
			time.Sleep(time.Second)
			j, more := <-jobs
			if more {
				fmt.Println("Received job", j)
			} else {
				fmt.Println("Received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 99; j >= 90; j-- {
		jobs <- j
		fmt.Println("Send job")
	}

	close(jobs)
	fmt.Println("Sent all jobs")

	<-done
}
