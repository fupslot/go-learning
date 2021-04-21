// reads the content of a file and prints its content to the stdout
// Notes:
//   - The file nto open should be provider as an argument os.Args
//   -  use io.Copy function
//   - "go run main.go in.txt"
package main

import (
	"fmt"
	"io"
	"os"
)

var usageStr = `

Usage "go run main.go in.txt"
`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usageStr)
		os.Exit(1)
	}

	fileName := os.Args[1]

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	defer f.Close()

	io.Copy(os.Stdin, f)
}
