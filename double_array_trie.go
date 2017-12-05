package go_tries

import (
	"strings"
	"bytes"
)

const (
	// Specifies an empty or available slot in the BC array
	//emptyValue = 0
	baseValue = 1
	// Rune to use a boundary between words
	boundary = "#"
	// Minimum numerical code
	minCode = 1
	// Maximum numerical code
	maxCode = 255
	// Slice grow number
	growInc = 16
)

type DoubleArrayTrie struct {
	// Base and check arrays
	base []int
	check []int
	// Tail array
	tail string
	// Current tail pos
	tailPos int
}

// Returns the current value of base
func (d *DoubleArrayTrie) getBase(pos int) int {
	idx := pos - 1
	if idx >= cap(d.base) {
		return 0
	}
	return d.base[idx]
}

// Returns the current value of check
func (d *DoubleArrayTrie) getCheck(pos int) int {
	idx := pos - 1
	if idx >= cap(d.check) {
		return 0
	}
	return d.check[idx]
}

func (d *DoubleArrayTrie) setBase(pos int, node int) {
	d.base = EnsureIndex(d.base, pos)
	d.base[pos - 1] = node
}

func (d *DoubleArrayTrie) setCheck(pos int, node int) {
	d.check = EnsureIndex(d.check, pos)
	d.check[pos - 1] = node
}

// Read tail starting at pos and ending in a boundary rune
func (d *DoubleArrayTrie) ReadTail(pos int) string {
	i := strings.Index(d.tail[pos:], boundary)
	if i != -1 && pos != i {
		return d.tail[:i+1]
	} else {
		return ""
	}
}

// Write at tail a text string starting at pos
func (d *DoubleArrayTrie) WriteTail(text string, pos int) {
	if pos == 0 {
		panic("Unexpected position parameter for WriteTail")
	}
	d.tail = d.tail[:pos-1]
	d.tail = d.tail + text
	d.tailPos = len(d.tail) + 1
}

// NewDoubleArrayTrie allocates and returns a new *DoubleArrayTrie.
func NewDoubleArrayTrie() *DoubleArrayTrie {
	d := &DoubleArrayTrie{
		base:   make([]int, 10, 10),
		check:   make([]int, 10, 10),
		tail: "",
		tailPos: 1,
	}
	// Set initial value of base at root
	d.setBase(1, baseValue)

	return d
}

func (d *DoubleArrayTrie) Get(key string) bool {
	idx, t := d.findTailPos(key)
	if idx == -1 {
		return false
	}
	// We still have to read the rest from the tail
	// compare it with the rest of the string
	if idx < len(key) {
		rest := d.ReadTail(-d.getBase(t))

		if rest == key[idx+1:] {
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

func (d *DoubleArrayTrie) Delete(key string) bool {
	idx, t := d.findTailPos(key)

	if idx == -1 {
		return false
	}

	// We still have to read the rest from the tail
	// compare it with the rest of the string
	if idx < len(key) {
		rest := d.ReadTail(-d.getBase(t))

		if rest == key[idx+1:] {
			// Clear out base and check
			d.setBase(t, 0)
			d.setCheck(t, 0)
			return true
		} else {
			return false
		}
	}

	return false
}

// Add specified key into trie. This method is similar to findTailPos
func (d *DoubleArrayTrie) Add(key string) bool {
	idx := -1
	s := 1
	var t int

	for {
		idx += 1

		ch := ValueFromChar(int(key[idx]))
		t = d.getBase(s) + ch

		// Case when check does not match with base. We have no match.
		if d.getCheck(t) != s {
			// Case when we have a conflict and we have to relocate the base
			if d.getCheck(t) != 0 {
				//d.relocateBase(s, t, key, idx)
			} else {
				// Case 1. Empty string or without conflicts. Just insert at tail
				d.separate(key, idx, s, d.tailPos)
				}
			return true
		}

		// Case when base denotes that the rest of the string
		// needs to be matched with the tail at pos
		if d.getBase(t) < 0 {
			break
		}

		// next word index
		s = t
	}

	// We still have to read the rest from the tail
	// compare it with the rest of the string. If match is found then the key is already inserted
	if idx < len(key) {
		rest := d.ReadTail(-d.getBase(t))

		if rest == key[idx+1:] {
			return true
		}
	}

	// We reached the end of the string, check if we have compared all chars,
	// otherwise attempt to insert at the end
	if idx == len(key) {
		return true
	} else {
		if d.getBase(t) != 0 {
			d.tailInsert(t, key)
		}
	}

	return false

}

// Update base and check by separating the first char of slice
func (d *DoubleArrayTrie) separate(slice string, idx int, s int, tailPos int) {
	checkPos := d.getBase(s) + ValueFromChar(int(slice[idx]))

	d.setCheck(checkPos, s)
	d.setBase(checkPos, -tailPos)
	d.WriteTail(slice[idx + 1:] + boundary, tailPos)
}


// Insert key into tail starting at tailPos
func (d *DoubleArrayTrie) tailInsert(s int, key string)  {
	// Save old pos
	oldTailPos := -d.getBase(s)

	// Init variables
	var list = []int{0, 0}
	idx := 0
	length := 0
	t := 0

	// Find length of common chars in tail
	for {
		if length > len(key) {
			break
		}
		// Find longest common prefix length between tail and key
		if key[length] == d.tail[length] {
			length += 1
		} else {
			break
		}
	}

	// Appends a sequence of arcs for the longest prefix
	for {
		// We have reached the end. Break now.
		if idx > length {
			break
		}
		// For each different character
		ch := d.tail[idx]

		list[0] = int(ch)
		// find next available place for common conflict at ch
		d.setBase(s, d.xCheck(list))
		// calculate next check origin
		t -= d.getBase(s + list[0])
		// Update check to point to base that was originated from
		d.setCheck(t, s)
		s = t

		idx += 1
	}

	list[0] = int(d.tail[length])
	list[1] = int(key[length])
	d.setBase(s, d.xCheck(list))

	d.separate(d.tail, length,s ,oldTailPos)
	d.separate(key, length, s, d.tailPos)
}

// Find max consecutive entries such as
// CHECK(BASE(s) + i) == s
func (d *DoubleArrayTrie) findArcs(s int) string {
	listLength := maxCode - minCode + 1
	var result bytes.Buffer
	i := minCode - 1
	var t int

	for {
		if i > listLength {
			break
		}
		i += 1
		t = d.getBase(s) + i
		c := d.getCheck(t)
		if c == s {
			result.WriteString(string(ValueToChar(i)))
		}
	}

	return result.String()
}

// Find minimum available q number such as CHECK(basePos + list[c]) !== 0
func (d *DoubleArrayTrie) xCheck(list []int) int {
	basePos := 1

	for {
		found := false

		for ch := 0; ch < len(list) ; ch += 1 {
			checkPos := d.getCheck(basePos + list[ch])

			if checkPos > 0 {
				found = true
				break
			}
		}

		if !found {
			break
		}

		basePos += 1
	}

	return basePos
}


func (d *DoubleArrayTrie) findTailPos(key string) (int, int) {
	idx := -1
	s := 1
	var t int

	for {
		idx += 1

		ch := ValueFromChar(int(key[idx]))
		t = d.getBase(s) + ch

		// Case when check does not match with base. We have no match.
		if d.getCheck(t) != s {
			return -1, -1
		}

		// Case when base denotes that the rest of the string
		// needs to be matched with the tail at pos
		if d.getBase(t) < 0 {
			break
		}

		// next word index
		s = t
	}

	return idx, t
}