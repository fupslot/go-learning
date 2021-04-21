package main

import (
	"fmt"
	"net/http"
	"time"
)

type Pipe string

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.yahoo.com",
	}

	ch := make(chan string)

	for _, link := range links {
		go checkStatus(link, ch)
	}

	for l := range ch {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkStatus(link, ch)
		}(l)
	}
}

func checkStatus(link string, ch chan string) {
	res, err := http.Head(link)
	if err != nil {
		fmt.Println(link, " seems down :(")
		ch <- link
		return
	}

	clen := res.Header.Get("content-length")
	if len(clen) == 0 {
		clen = "0"
	}

	fmt.Printf("\t%s\t(%s bytes)\tup!\n", link, clen)
	ch <- link
}
