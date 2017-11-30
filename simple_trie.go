package go_tries

import (
	"strings"
)

type SimpleTrie struct {
	// Reference to children
	children map[string]*SimpleTrie
	// Value of Node
	value interface{}
}

// Get Next word from a key, a starting index and a path separator
func NextWord(key string, start int, sep rune) (segment string, nextIndex int) {
	if len(key) == 0 || start < 0 || start > len(key)-1 {
		return "", -1
	}
	end := strings.IndexRune(key[start:], sep) // next sep after 0th rune
	if end == -1 {
		return key[start:], -1
	}
	return key[start: start+end+1], start + end + 1
}


// NewSimpleTrie allocates and returns a new *SimpleTrie.
func NewSimpleTrie() *SimpleTrie {
	return &SimpleTrie{
		children: make(map[string]*SimpleTrie),

	}
}

// Get returns the value stored at the given key. Returns nil for internal
// nodes or for nodes with a value of nil.
func (trie *SimpleTrie) Get(key string) interface{} {
	node := trie
	for part, i := NextWord(key, 0, ' '); ; part, i = NextWord(key, i, ' ') {
		node = node.children[part]
		if node == nil {
			return nil
		}
		if i == -1 {
			break
		}

	}
	return node.value
}

func (trie *SimpleTrie) Add(key string, value int) bool {
	node := trie
	for part, i := NextWord(key, 0, ' '); ; part, i = NextWord(key, i, ' ') {

		child, _ := node.children[part]

		if child == nil {
			child = NewSimpleTrie()
			node.children[part] = child
		}

		node = child
		if i == -1 {
			break
		}

	}

	isNewVal := node.value == nil
	node.value = value

	return isNewVal
}

// PathTrie node and the part string key of the child the path descends into.
type nodeStr struct {
	node *SimpleTrie
	part string
}

func (trie *SimpleTrie) Delete(key string) bool {
	var path []nodeStr // record ancestors to check later
	node := trie
	for part, i := NextWord(key, 0, ' '); ; part, i = NextWord(key, i, ' ') {
		path = append(path, nodeStr{part: part, node: node})
		node = node.children[part]
		if node == nil {
			// node does not exist
			return false
		}
		if i == -1 {
			break
		}
	}

	// delete the node value
	node.value = nil

	// if leaf, remove it from its parent's children map. Repeat for ancestor path.
	if len(node.children) == 0 {
		// iterate backwards over path
		for i := len(path) - 1; i >= 0; i-- {
			parent := path[i].node
			part := path[i].part
			delete(parent.children, part)
			if parent.value != nil || !(len(parent.children) == 0) {
				// parent has a value or has other children, stop
				break
			}
		}
	}

	return true
}
