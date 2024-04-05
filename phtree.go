package phtree

import (
	"math/bits"
	"sync"
)

// PHTree represents a Permutation Hierarchical Tree.
type PHTree struct {
	root     *node
	maxDepth int
	mu       sync.RWMutex
}

// node represents a node in the PHTree.
type node struct {
	children []*node
	entries  []entry
	mu       sync.RWMutex
}

// entry represents a key-value pair in the PHTree.
type entry struct {
	key   []uint64
	value interface{}
}

// Option represents an option for configuring the PHTree.
type Option func(*PHTree)

// WithMaxDepth sets the maximum depth of the PHTree.
func WithMaxDepth(maxDepth int) Option {
	return func(t *PHTree) {
		t.maxDepth = maxDepth
	}
}

// New creates a new PHTree with the specified options.
func New(opts ...Option) *PHTree {
	t := &PHTree{
		root: &node{
			children: make([]*node, 2),
		},
		maxDepth: 64,
	}

	for _, opt := range opts {
		opt(t)
	}

	return t
}

// Insert inserts a key-value pair into the PHTree.
func (t *PHTree) Insert(key []uint64, value interface{}) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.root.insert(key, value, 0, t.maxDepth)
}

// Search searches for a key in the PHTree and returns the associated value.
func (t *PHTree) Search(key []uint64) interface{} {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.root.search(key, 0, t.maxDepth)
}

// Remove removes a key-value pair from the PHTree.
func (t *PHTree) Remove(key []uint64) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.root.remove(key, 0, t.maxDepth)
}

// insert inserts a key-value pair into the node or its children.
func (n *node) insert(key []uint64, value interface{}, depth, maxDepth int) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if depth == len(key) || depth == maxDepth {
		n.entries = append(n.entries, entry{key, value})
		return
	}

	bit := key[depth] & 1
	if n.children[bit] == nil {
		n.children[bit] = &node{
			children: make([]*node, 2),
		}
	}
	n.children[bit].insert(key, value, depth+1, maxDepth)
}

// search searches for a key in the node or its children and returns the associated value.
func (n *node) search(key []uint64, depth, maxDepth int) interface{} {
	n.mu.RLock()
	defer n.mu.RUnlock()

	if depth == len(key) || depth == maxDepth {
		for _, e := range n.entries {
			if equalKeys(e.key, key) {
				return e.value
			}
		}
		return nil
	}

	bit := key[depth] & 1
	if n.children[bit] != nil {
		return n.children[bit].search(key, depth+1, maxDepth)
	}
	return nil
}

// remove removes a key-value pair from the node or its children.
func (n *node) remove(key []uint64, depth, maxDepth int) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if depth == len(key) || depth == maxDepth {
		for i, e := range n.entries {
			if equalKeys(e.key, key) {
				n.entries = append(n.entries[:i], n.entries[i+1:]...)
				break
			}
		}
		return
	}

	bit := key[depth] & 1
	if n.children[bit] != nil {
		n.children[bit].remove(key, depth+1, maxDepth)
		if len(n.children[bit].entries) == 0 && n.children[bit].isLeaf() {
			n.children[bit] = nil
		}
	}
}

// isLeaf checks if the node is a leaf node.
func (n *node) isLeaf() bool {
	return n.children[0] == nil && n.children[1] == nil
}

// equalKeys checks if two keys are equal.
func equalKeys(a, b []uint64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// bitCount counts the number of set bits in a uint64.
func bitCount(x uint64) int {
	return bits.OnesCount64(x)
}
