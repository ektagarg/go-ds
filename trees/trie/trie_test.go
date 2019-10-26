package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var trie = CreateTrie()

func TestInsert(t *testing.T) {
	trie.Insert(10, []rune("a"))
	trie.Insert(11, []rune("b"))
	trie.Insert(12, []rune("c"))

	trie.Insert(13, []rune("and"))
	trie.Insert(14, []rune("as"))
	trie.Insert(15, []rune("assign"))
	trie.Insert(16, []rune("ah don't know"))
}

func TestDelete(t *testing.T) {
	// ...
}

func TestSearch(t *testing.T) {
	lookup := []string{"a", "b", "c", "and", "as", "assign", "ah don't know"}

	expected := []Data{10, 11, 12, 13, 14, 15, 16}

	for i, look := range lookup {
		ans, err := trie.Search([]rune(look))
		assert.Equal(t, ans, expected[i], "should be equal")
		assert.Equal(t, err, nil, "should be equal")
	}
}
