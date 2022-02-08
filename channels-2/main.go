package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var ch = make(chan string)
	var workers = []int64{1, 2}

	var wg sync.WaitGroup
	for _, w := range workers {
		wg.Add(1)
		go func(wn string) {
			var ticker *time.Ticker = time.NewTicker(time.Second)

			fmt.Printf("%s started\n", wn)
			defer func() {
				wg.Done()
				fmt.Printf("%s end\n", wn)
			}()

			c := 0
			for t := range ticker.C {
				if c >= 10 {
					return
				}
				ch <- fmt.Sprintf("%s -> %d", wn, t.Unix())
				c += 1
			}
		}(fmt.Sprintf("worker #%d", w))
		time.Sleep(time.Second)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("%s\n", i)
	}

	fmt.Println("done")
}
