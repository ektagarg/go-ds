package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var trie = CreateTrie()
var insertEntries = []string{
	"and",
	"a",
	"b",
	"c",
	"andromeda",
	"as",
	"assign",
	"ah don't know",
}

var deleteEntries = []string{
	"and",
	"assign",
}

func TestInsert(t *testing.T) {
	for i, entry := range insertEntries {
		trie.Insert(Data(i), []rune(entry))
	}
	trie.PrintTrie()
}

func TestSearch(t *testing.T) {
	for expected, lookup := range insertEntries {
		ans, err := trie.Search([]rune(lookup))
		assert.Equal(t, Data(expected), ans, "should be equal")
		assert.Equal(t, nil, err, "should be equal")
	}
}

func TestDelete(t *testing.T) {
	for _, entry := range deleteEntries {
		trie.Delete([]rune(entry))
	}
	trie.PrintTrie()
}
