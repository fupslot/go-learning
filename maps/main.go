package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Name  string
	Likes []string
}

func main() {
	var people []*Person

	people = append(people, &Person{Name: "John Doe", Likes: []string{"cheese", "cat"}})
	people = append(people, &Person{Name: "Alice Wonder", Likes: []string{"sleep"}})
	people = append(people, &Person{Name: "Colin Doe", Likes: []string{"cheese", "dog"}})
	people = append(people, &Person{Name: "Melissa May", Likes: []string{"sleep", "cat"}})

	likes := make(map[string][]*Person)

	for _, p := range people {
		for _, l := range p.Likes {
			likes[l] = append(likes[l], p)
		}
	}

	// Print people that like cat
	for _, p := range likes["cat"] {
		fmt.Println(p.Name, "likes cat!")
	}

	// Ptin numnber of people that like cat
	fmt.Println(len(likes["cat"]), "people like cat!")

	type Key struct {
		Path, Country string
	}

	hits := make(map[Key]int)

	hits[Key{"/doc/", "au"}]++

	fmt.Println("/doc/", hits[Key{"/doc/", "au"}])

	// Maps are not safe for concurrent use:
	// it's not defined what happens when you read and write to them simultaneously.

	// protect maps is with sync.RWMutex.

	var m = struct {
		sync.RWMutex
		hits map[Key]int
	}{hits: make(map[Key]int)}

	m.Lock()
	m.hits[Key{"/doc/", "EU"}]++
	m.Unlock()

	m.RLock()
	fmt.Println(m.hits[Key{"/doc/", "EU"}])
	m.RUnlock()

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	changeMap(colors)

	printMap(colors)
}

func changeMap(c map[string]string) {
	c["yellow"] = "#FFFF00"
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("color %s\thex %s\n", color, hex)
	}
}
