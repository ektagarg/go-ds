package trie

import "testing"

var trie = CreateTrie()

func TestInsert(t *testing.T) {
	trie.Insert(10, []rune("a"))
	trie.Insert(11, []rune("b"))
	trie.Insert(12, []rune("c"))
}

func TestDelete(t *testing.T) {
	// ...
}

func TestSearch(t *testing.T) {
	// ...
}
