package tree

import (
	"errors"
	"strings"
)

// Trie is the root of the tree.
// It doesn't have any data or prefixes.
type Trie struct {
	Children []*Node
}

// Node is a node of the Trie.
type Node struct {
	Data     string
	Prefix   string
	Children []*Node
}

// CreateTrie creates an empty Trie tree and returns it.
func CreateTrie() Trie {
	var t Trie
	return t
}

// CreateNode creates a new node containing data and prefix string.
// Returns a pointer to that node.
func CreateNode(data, prefix string) *Node {
	return &Node{data, prefix, nil}
}

// Insert appends a node (n) containing data and prefix to the trie.
// Returns an error if the node already exists.
func (t Trie) Insert(data, prefix string) error {

	if prefix == "" || prefix == nil {
		return errors.New("Can't insert node with empty prefix")
	}

	if t == nil {
		return errors.New("Can't insert on nil trie")
	}

	n := CreateNode(data, prefix)

	return t.insertNode(n)
}

func insertNode(n *Node) error {
}

// Delete searches for a prefix in the Trie.
// Removes the node and rearranges the tree if prefix exists.
// Returns an error if prefix doesn't exist.
func (t Trie) Delete(prefix string) error {
}

// Search looks for the node indexed by prefix.
// Returns a string containing the data if prefix exists.
// Returns an empty string and error if prefix doesn't exist.
func (t Trie) Search(prefix string) (string, error) {
	if t == nil {
		return "", errors.New("Can't search in nil trie")
	}

	for aux := range t.Children {
		if strings.HasPrefix(prefix, aux.Prefix) {
			return searchFromNode(prefix, aux)
		}
	}

	return "", errors.New("Didn't find prefix")
}

func searchFromNode(prefix string, node *Node) (string, error) {
	// Found prefix on current node
	if node.Prefix == prefix {
		return node.Data, nil
	}

	for aux := range node.Children {
		if strings.HasPrefix(prefix, aux.Prefix) {
			return searchFromNode(prefix, aux)
		}
	}

	// Couldn't find prefix (node doesn't exist)
	return nil, errors.New("Couldn't find prefix")
}
