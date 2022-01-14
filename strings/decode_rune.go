package main

import (
	"fmt"
	"io/ioutil"
	"unicode/utf8"
	"os"
)

func main() {
	content, err := ioutil.ReadFile("text.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for len(content) > 0 {
		r, size := utf8.DecodeRune(content)
		fmt.Printf("%c\t%v\n", r, size)
		content = content[size:]
	}
}
