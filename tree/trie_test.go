package trie

import "testing"

var trie = CreateTrie()

func TestInsert(t *testing.T) {
	trie.Insert(10, "a")
	trie.Insert(11, "b")
	trie.Insert(12, "c")
}

func TestDelete(t *testing.T) {
	// ...
}

func TestSearch(t *testing.T) {
	// ...
}
