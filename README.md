# Go Tries


Go tries implements several types of Tries for experimental use.

The implementations are optimized for Get performance and to allocate 
0 bytes of heap memory (i.e. garbage) per Get.

A typical use case is to perform Puts and Deletes upfront to populate the Trie, 
then perform Gets very quickly.

When Tries are chosen over maps, it is typically for their space efficiency.


Trie Types
---

**SimpleTrie**: A simple implementation using a map of TrieNodes.

```go
t := NewSimpleTrie()
t.Add("cat", 0)
t.Add("fox", 1)
t.Add("dog", 2)
t.Add("dog and", 3)
t.Add("dog and cat", 4)

t.Get("Cat") // nil
t.Get("cat") // cat
```

**DoubleArrayTrie**: A more complex implementation of a Trie using 2 Lists. 
This is supposed to have better search performance in expense of slower insertions.
Based on the [Sato and Morimoto paper](http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.14.8665&rep=rep1&type=pdf)

Benchmarks
---
Single threaded benchmarks: Simple Trie
```bash
BenchmarkSimpleTriePutStringKey-4       50000000                35.6 ns/op             8 B/op          1 allocs/op
BenchmarkSimpleTrieGetStringKey-4       100000000               16.0 ns/op             0 B/op          0 allocs/op
BenchmarkSimpleTriePutPhraseKey-4       20000000                69.5 ns/op             8 B/op          1 allocs/op
BenchmarkSimpleTrieGetPhraseKey-4       30000000                42.1 ns/op             0 B/op          0 allocs/op
```

Single threaded benchmarks: Double Array Trie
```bash
BenchmarkDoubleArrayTrieGetSimpleStringKey-4    50000000                26.7 ns/op             0 B/op          0 allocs/op
```

License
---

MIT License