package tree

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

// Insert adds a node (n) to the Trie.
// Returns an error if the node already exists.
func (t Trie) Insert(n Node) error {
}

// Delete searches for a prefix in the Trie.
// Removes the node and rearranges the tree if prefix exists.
// Returns an error if prefix doesn't exist.
func (t Trie) Delete(prefix string) error {
}

// Search looks for the node indexed by prefix.
// Returns a string containing the data if prefix exists.
// Returns an error if prefix doesn't exist.
func (t Trie) Search(prefix string) (string, error) {
}
