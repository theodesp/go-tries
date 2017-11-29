package go_tries


import (
	"crypto/rand"
	"testing"
)

var words [1000]string // random string words
const bytesPerKey = 30

var phrases [1000]string // random phrases
const wordsPerPhrase = 3     // (e.g. a b c has parts a, b, c)
const bytesPerPart = 10

func init() {
	// word keys
	for i := 0; i < len(words); i++ {
		key := make([]byte, bytesPerKey)
		if _, err := rand.Read(key); err != nil {
			panic("error generating random byte slice")
		}
		phrases[i] = string(key)
	}

	// path keys
	for i := 0; i < len(phrases); i++ {
		var key string
		for i := 0; i < wordsPerPhrase; i++ {
			key += " "
			part := make([]byte, bytesPerPart)
			if _, err := rand.Read(part); err != nil {
				panic("error generating random byte slice")
			}
			key += string(part)
		}
		phrases[i] = string(key)
	}
}

func TestNilCases(t *testing.T)  {
	b := NewSimpleTrie()

	cases := []struct {
		key   string
		value int
	}{
		{"fish", 0},
		{"cat", 1},
		{"dog", 2},
		{"cats", 3},
		{"caterpillar", 4},
		{"cat gideon", 5},
		{"cat giddy", 6},
		{"cat maker", 7},
	}

	// subsequent put
	for _, c := range cases {
		b.Add(c.key, c.value)
	}

	expectNilValues := []string{"", "c", "ca", "caterpillar2", "other"}

	// get nil
	for _, key := range expectNilValues {
		if value := b.Get(key); value != nil {
			t.Errorf("expected key %s to have value nil, got %v", key, value)
		}
	}
}


func BenchmarkSimpleTriePutStringKey(b *testing.B) {
	trie := NewSimpleTrie()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		trie.Add(words[i%len(words)], i)
	}
}

func BenchmarkSimpleTrieGetStringKey(b *testing.B) {
	trie := NewSimpleTrie()
	for i := 0; i < b.N; i++ {
		trie.Add(words[i%len(words)], i)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		trie.Get(words[i%len(words)])
	}
}

// Phrase keys

func BenchmarkSimpleTriePutPhraseKey(b *testing.B) {
	trie := NewSimpleTrie()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		trie.Add(phrases[i%len(phrases)], i)
	}
}

func BenchmarkSimpleTrieGetPhraseKey(b *testing.B) {
	trie := NewSimpleTrie()
	for i := 0; i < b.N; i++ {
		trie.Add(phrases[i%len(phrases)], i)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		trie.Get(phrases[i%len(phrases)])
	}
}
