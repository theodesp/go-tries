package go_tries

import (
	"strings"
	"fmt"
)

const (
	// Specifies an empty or available slot in the BC array
	//emptyValue = 0
	baseValue = 1
	// Rune to use a boundary between words
	boundary = "#"
	//// Minimum numerical code
	//minCode = 1
	//// Maximum numerical code
	//maxCode = 255
	//// Slice grow number
	growInc = 16
)

type DoubleArrayTrie struct {
	// Base and check arrays
	bc []int
	// Tail array
	tail string
	// Current tail pos
	tailPos int
}

// Returns the current value of base
func (d *DoubleArrayTrie) getBase(pos int) int  {
	idx := 2 * pos
	if idx > cap(d.bc) {
		return 0
	}
	return d.bc[idx]
}

// Returns the current value of check
func (d *DoubleArrayTrie) getCheck(pos int) int  {
	idx := 2 * pos + 1
	if idx > cap(d.bc) {
		return 0
	}
	return d.bc[idx]
}


func (d *DoubleArrayTrie) setBase(pos int, node int)  {
	d.bc = EnsureIndex(d.bc, 2 * pos)
	d.bc[2 * pos] = node
}

func (d *DoubleArrayTrie) setCheck(pos int, node int)  {
	d.bc = EnsureIndex(d.bc, 2 * pos + 1)
	d.bc[2 * pos + 1] = node
}

// Read tail starting at pos and ending in a boundary rune
func (d *DoubleArrayTrie) ReadTail(pos int) string  {
	i := strings.Index(d.tail[pos:], boundary)
	if i != -1 && pos != i {
		return d.tail[:i+1]
	} else {
		return ""
	}
}

// Write at tail a text string starting at pos
func (d *DoubleArrayTrie) WriteTail(text string, pos int)  {
	d.tail = d.tail[:pos]
	d.tail = d.tail + text
	d.tailPos = len(d.tail)
}

// NewDoubleArrayTrie allocates and returns a new *DoubleArrayTrie.
func NewDoubleArrayTrie() *DoubleArrayTrie {
	d := &DoubleArrayTrie{
		bc: make([]int, 10, 10),
		tail: boundary,
	}
	// Set initial value of base at root
	d.setBase(1, baseValue)

	return d
}

func (d *DoubleArrayTrie) Get(key string) bool  {
	idx := -1
	s := 1
	var t int

	for {
		idx += 1

		ch := int(key[idx]) - 'a' + 1
		t = d.getBase(s) + ch

		// Case when check does not match with base. We have no match.
		if d.getCheck(t) != s {
			return false
		}

		// Case when base denotes that the rest of the string
		// needs to be matched with the tail at pos
		if d.getBase(t) < 0 {
			break
		}

		// next word index
		s=t
	}

	// We still have to read the rest from the tail
	// compare it with the rest of the string
	if idx < len(key) {
		rest := d.ReadTail(-d.getBase(t))

		if rest == key[idx:] {
			return true
		} else {
			return false
		}
	}

	// We reached the end of the string, check if we have compared all chars,
	// otherwise fail as we still have chars left in tail
	if idx == len(key) {
		return true
	} else {
		return false
	}
}

func (d *DoubleArrayTrie) toString() string  {
	return fmt.Sprintf("%+v", d)
}
