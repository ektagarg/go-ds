package tree

import (
	"errors"
	"strings"
)

// Data is the data type the trie holds

type Data int

// Trie is the root of the tree.
// It doesn't have any data or prefixes.
type Trie struct {
	root *Node
}

// Node is a node of the Trie.
type Node struct {
	data       Data
	isTerminal bool
	prefix     rune
	children   []*Node
}

// CreateTrie creates an empty Trie tree and returns it.
func CreateTrie() *Trie {
	return &Trie{&Node{nil, nil, nil, nil}}
}

// Returns a pointer to that node.
// CreateNode creates a new node.
func CreateNode(data Data, isTerminal bool, prefix rune) *Node {
	return &Node{data, isTerminal, prefix, nil}
}

// Insert appends a node (n) containing data and prefix to the trie.
// Returns an error if the node already exists.
func (t *Trie) Insert(data Data, prefix string) error {

	if prefix == "" || prefix == nil {
		return errors.New("Can't insert node with empty prefix")
	}

	if t == nil {
		return errors.New("Can't insert on nil trie")
	}

	n := t.root
	// Move to the last existing node
	for c := range prefix {
		if !n.charInNodeChildren(c) { // Go to next child
			// createSubTree(n, rest_of_prefix)
			break // Stops when child with prefix char doesn't exist
		}
	}

	n := CreateNode(data, prefix)

	return t.insertNode(n)
}

// charInNodeChildren returns true if it finds c as
// a prefix in any of n's children and points n to
// the child where it found c.
// Returns false and n remains unaltered otherwise
func (n *Node) charInNodeChildren(c rune) bool {

	if n.children == nil {
		return false
	}

	for child := range n.children {
		if child.prefix == c {
			n = child
			return true
		}
	}

	return false
}

func (t *Trie) insertNode(n *Node) error {

}

// Delete searches for a prefix in the Trie.
// Removes the node and rearranges the tree if prefix exists.
// Returns an error if prefix doesn't exist.
func (t *Trie) Delete(prefix string) error {
}

// Search looks for the node indexed by prefix.
// Returns a string containing the data if prefix exists.
// Returns an empty string and error if prefix doesn't exist.
func (t *Trie) Search(prefix string) (string, error) {
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
