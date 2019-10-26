package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var trie = CreateTrie()
var entries = []string{
	"and",
	"a",
	"b",
	"c",
	"andromeda",
	"as",
	"assign",
	"ah don't know",
}

func TestInsert(t *testing.T) {
	for i, entry := range entries {
		trie.Insert(Data(i), []rune(entry))
	}
	trie.PrintTrie()
}

func TestDelete(t *testing.T) {
	// ...
}

func TestSearch(t *testing.T) {
	for expected, lookup := range entries {
		ans, err := trie.Search([]rune(lookup))
		assert.Equal(t, Data(expected), ans, "should be equal")
		assert.Equal(t, nil, err, "should be equal")
	}
}
