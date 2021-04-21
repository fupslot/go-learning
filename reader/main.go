package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var (
		InputString string
	)

	flag.StringVar(&InputString, "input", "", "Input string to reverse")
	flag.Parse()

	fmt.Printf("%s", InputString)

	if InputString == "" {
		ib, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Error", err)
			os.Exit(2)
		}

		InputString = string(ib)
	}

	if InputString == "" {
		printUsage()
		os.Exit(1)
	}

	st := []byte(InputString)

	l := len(st)

	/**
	* The idea is to iterate over the half of a string and exchange
	* characters from one end to another.
	* "stash" variable used to store a character of an iteration
	*
	*       stash
	*    /  /    ^  ^
	*   /  /      \  \
	* h | e | l | l | o
	 */

	// Devide the string lenght on half (sort of)
	it := l >> 1 // faster devision ?

	fmt.Printf("Iterations: %d\n", it)

	var stash byte

	for i := it - 1; i >= 0; i-- {
		stash = st[i]     // stashing
		st[i] = st[l-1-i] // last become first
		st[l-1-i] = stash // first become last
	}

	fmt.Printf("Result: %s", st)
}

var usageStr = `
Usage: main [options]

Options:
        --input <string> Input string to reverse
`

func printUsage() {
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}
