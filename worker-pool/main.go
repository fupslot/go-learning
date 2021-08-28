package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan Job, result chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %s \n", id, j.Name)
		// https://stackoverflow.com/questions/17573190/how-to-multiply-duration-by-integer
		time.Sleep(time.Duration(int32(j.Value*2) * int32(time.Second)))
		fmt.Printf("Worker %d finished job %s \n", id, j.Name)

		result <- j.Value * 2
	}
}

type Job struct {
	Name  string
	Value int
}

func main() {
	const jobs_total = 5

	jobs := make(chan Job, jobs_total)
	results := make(chan int, jobs_total)

	// start up 3 workes
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// send 5 jobs
	for j := 1; j <= jobs_total; j++ {
		jobs <- Job{Name: fmt.Sprintf("JOB_%d", j), Value: j}
	}

	nums := []int{}

	for r := 1; r <= jobs_total; r++ {
		nums = append(nums, <-results)
	}

	fmt.Println("Results", nums)
}
