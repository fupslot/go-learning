package main

import (
	"fmt"
	"os"
	"strings"
	"text/scanner"
	"unicode/utf8"
)

func main() {
	fmt.Println("Contains: ", strings.Contains("test", "st"))
	fmt.Println("Count:", strings.Count("test", "t"))

	var input string = "hello \tworld!\n"
	whitespace := "\t\n "
	fmt.Println(input[3:])

	for len(input) > 0 {
		r, s := utf8.DecodeRuneInString(input)
		// fmt.Printf("%c\t%v\n", r, s)
		input = input[s:]

		i := strings.IndexRune(whitespace, r)
		fmt.Printf("%c\t%v (%d)\n", r, s, i)
	}

	f, _ := os.Open("file")
	defer f.Close()

	var sc scanner.Scanner
	scan := sc.Init(f)

	for {
		r := scan.Next()
		if r == scanner.EOF {
			return
		}
		fmt.Printf("%d\t%c\n", r, r)
	}
}
