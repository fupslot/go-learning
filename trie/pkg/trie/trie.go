package trie

import (
	"unicode/utf8"
)

type Node struct {
	Children map[rune]*Node
	EOW      bool
}

type Trie struct {
	Root *Node
	sug  []string
	lim  int32
}

func CreateTrie() *Trie {
	return &Trie{
		Root: &Node{
			Children: make(map[rune]*Node),
		},
		lim: 2,
	}
}

func (t *Trie) Insert(word string) {
	node := t.Root
	for i := 0; i < len(word); i++ {
		r, _ := utf8.DecodeRuneInString(word[i:])

		if _, ok := node.Children[r]; !ok {
			node.Children[r] = &Node{
				Children: make(map[rune]*Node),
			}
		}

		node = node.Children[r]
	}
	node.EOW = true
}

func (t *Trie) Search(key string) bool {
	node := t.Root

	for i := 0; i < len(key); i++ {
		r, _ := utf8.DecodeRuneInString(key[i:])
		if _, ok := node.Children[r]; !ok {
			return false
		}
		node = node.Children[r]
	}

	return true
}

func (t *Trie) traverse(node *Node, word string) {
	if node.EOW {
		t.sug = append(t.sug, word)
	}

	for r, node := range node.Children {
		if len(t.sug) == int(t.lim) { // limiting number suggestions
			break
		}
		t.traverse(node, word+string(r))
	}
}

func (t *Trie) Autocomplete(key string) []string {
	exists := true
	node := t.Root

	t.sug = make([]string, 0)

	if len(key) == 0 {
		return t.sug
	}

	for i := 0; i < len(key); i++ {
		r, _ := utf8.DecodeRuneInString(key[i:])
		if _, ok := node.Children[r]; !ok {
			exists = false
			break
		}

		node = node.Children[r]
	}

	if !exists {
		return t.sug
	}

	t.traverse(node, key)

	return t.sug
}
