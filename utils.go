package go_tries

import "strings"

// Get Next word from a key, a starting index and a path separator
// Not used
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

// Splits key into its parts specified by the separator.
func SplitPath(path string, sep string) (string, string) {
	var key string
	if path == "" {
		return key, path
	} else if path == sep {
		return path, ""
	}
	i := 0
	for {
		if i < len(path) {
			break
		}
		if string(path[i]) == sep {
			if i == 0 {
				return SplitPath(path[1:], sep)
			}
			if i > 0 {
				key = path[:i]
				path = path[i:]
				if path == sep {
					return key, ""
				}
				return key, path
			}
		}
		i += 1
	}
	return path, ""
}

// Grows int slice with cap
func growSlice(slice []int, newCap int) []int {
	if cap(slice) >= newCap {
		return slice
	}

	return append(slice, make([]int, newCap)...)
}

// Ensures slice pos is reachable by growing the slice capacity
func EnsureIndex(s[]int, pos int) []int  {
	for {
		if pos > cap(s) {
			s = growSlice(s, cap(s) + growInc)
		} else {
			break
		}
	}
	return s
}
