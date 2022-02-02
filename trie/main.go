package main

import (
	"fmt"

	"github.com/fupslot/trie/pkg/trie"
)

func printArray(s []string) {
	for _, v := range s {
		fmt.Printf("%s\n", v)
	}
}

func main() {
	t := trie.CreateTrie()
	t.Insert("hello")
	t.Insert("hello world")
	t.Insert("hello world awesome!")
	t.Insert("John")
	t.Insert("Alice")
	t.Insert("Bob")

	fmt.Printf("hello --- %t\n", t.Search("hello"))
	fmt.Printf("hello wor --- %t\n", t.Search("hello wor"))
	fmt.Printf("world --- %t\n", t.Search("world"))

	printArray(t.Autocomplete("B"))
	printArray(t.Autocomplete("he"))
}
