package go_tries

// Abstract interface for Trie
type Trie interface {
	Get(key string) interface{}
	Add(key string, value int) bool
	Delete(key string) bool
}
