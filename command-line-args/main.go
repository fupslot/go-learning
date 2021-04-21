// Echo1 prints its command-line arguments.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	fmt.Print(files)

	if len(files) == 0 {
		countLines(os.Stdin, counts)
		printCounts(counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts)
			printCounts(counts)

			f.Close()
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func printCounts(counts map[string]int) {
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
