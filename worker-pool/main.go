package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j)
		// https://stackoverflow.com/questions/17573190/how-to-multiply-duration-by-integer
		time.Sleep(time.Duration(int32(j*2) * int32(time.Second)))
		fmt.Println("Worker", id, "finished job", j)

		result <- j * 2
	}
}

func main() {
	const jobs_total = 5

	jobs := make(chan int, jobs_total)
	results := make(chan int, jobs_total)

	// start up 3 workes
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// send 5 jobs
	for j := 1; j <= jobs_total; j++ {
		jobs <- j
	}

	nums := []int{}

	for r := 1; r <= jobs_total; r++ {
		nums = append(nums, <-results)
	}

	fmt.Println("Results", nums)
}
