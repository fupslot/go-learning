package main

import (
	"fmt"
	"os"
	"strings"
	"text/scanner"
	"unicode/utf8"
)

func Playground() {
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

var keypad_map map[rune]int = map[rune]int{
	'a': 2, 'b': 2, 'c': 2,
	'd': 3, 'e': 3, 'f': 3,
	'g': 4, 'h': 4, 'i': 4,
	'j': 5, 'k': 5, 'l': 5,
	'm': 6, 'n': 6, 'o': 6,
	'p': 7, 'q': 7, 'r': 7, 's': 7,
	't': 8, 'u': 8, 'v': 8,
	'w': 9, 'x': 9, 'y': 9, 'z': 9,
}

func toNumber(word string) []int {
	var n []int = make([]int, len(word))
	for i := 0; i < len(n); i++ {
		n[i] = keypad_map[rune(word[i])]
	}
	return n
}

func phoneNumber(letters string) []int {
	var numbers []int = make([]int, len(letters))
	for i := 0; i < len(numbers); i++ {
		numbers[i] = int(rune(letters[i]) - rune('0'))
	}

	return numbers
}

func main() {
	input := []string{"flower", "hello", "tigerking"}
	fmt.Printf("%v\n", toNumber("hello"))
	fmt.Printf("%v\n", phoneNumber("470133"))

	var words [][]int = make([][]int, len(input))
	for i := 0; i < len(words); i++ {
		words[i] = toNumber(input[i])
	}

	fmt.Printf("%v", words)

	// playground()
}
