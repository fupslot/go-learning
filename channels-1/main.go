package main

import (
	"fmt"
	"net/http"
)

func main() {
	depsFinished := make(chan bool)
	defer close(depsFinished)

	go prerequsites(depsFinished)

	if <-depsFinished {
		fmt.Println("Done")
	} else {
		fmt.Println("Something went wrong!!!")
	}
}

const depsURL = "https://raw.githubusercontent.com/fupslot/angular-hotkeys-light/master/package.json"

func prerequsites(done chan bool) {
	fmt.Println("Installing dependensies...")

	res, err := http.Get(depsURL)
	if err != nil {
		fmt.Println(err.Error())
		done <- false
		return
	}

	if res.StatusCode != 200 {
		done <- false
		return
	}

	// fmt.Printf("Status: %d\n", res.StatusCode)

	fmt.Println("All dependensies installed!")
	done <- true
}
