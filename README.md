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
b.Add("cat", 0)
b.Add("fox", 1)
b.Add("dog", 2)
b.Add("dog and", 3)
b.Add("dog and cat", 4)

b.Get("Cat") // false
b.Get("cat") // true
```

License
---

MIT License