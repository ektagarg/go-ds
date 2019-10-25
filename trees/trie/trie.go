package trie

import (
	"errors"
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
	for i, c := range prefix {
		if !n.charInNodeChildren(c) { // Go to next child
			lastChar := i
			break // Stops when child with prefix char doesn't exist
		}
	}

	remainingPrefix = []rune(prefix)[lastChar:]
	return n.createSubTree(data, remainingPrefix)
}

// createSubTree inserts the rest of a prefix beginning
// in the Node n.
func (n *Node) createSubTree(data Data, prefix []rune) error {
	var newNode *Node = nil
	var newNodeDad *Node = n

	for _, c := range prefix {
		newNode = CreateNode(data, false, c)
		newNodeDad.children = append(newNodeDad, newNode)
		newNodeDad = newNode
	}

	// Insert data in last node
	newNode.isTerminal = true
	newNode.data = data

	return nil
}

// charInNodeChildren returns true if it finds c as
// a prefix in any of n's children and points n to
// the child where it found c.
// Returns false and n remains unaltered otherwise
func (n *Node) charInNodeChildren(c rune) bool {

	if n.children == nil {
		return false
	}

	for _, child := range n.children {
		if child.prefix == c {
			n = child
			return true
		}
	}

	return false
}

// Delete searches for a prefix in the Trie.
// Removes the node and rearranges the tree if prefix exists.
// Returns an error if prefix doesn't exist.
func (t *Trie) Delete(prefix string) error {
}

// Search looks for the node indexed by prefix.
// Returns a string containing the data if prefix exists.
// Returns an empty string and error if prefix doesn't exist.
func (t *Trie) Search(prefix []rune) (Data, error) {
	if t == nil {
		return nil, errors.New("Can't search in nil trie")
	}

	if prefix == nil {
		return nil, errors.New("Can't search nil prefix")
	}

	// Lookup node starts at root
	lookup := t.root

	for _, c := range prefix {
		for i, n := range lookup.children {
			if n.prefix == c { // found prefix, update lookup (go down)
				lookup = n
				break
			}
			if i == len(lookup.children)-1 { // at the last child of lookup
				return nil, errors.New("Didn't find prefix")
			}
		}
	}

	if lookup.isTerminal {
		return lookup.data, nil
	}
	return nil, errors.New("Leaf node is not terminal")
}
