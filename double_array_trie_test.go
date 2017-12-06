package go_tries

import (
	"testing"
)

func TestInitTail(t *testing.T) {
	d := NewDoubleArrayTrie()

	if d.tail != "" {
		t.Errorf("expected tail initial value to be %v, got %v", "", d.tail)
	}
}

func TestInitBase(t *testing.T) {
	d := NewDoubleArrayTrie()

	if d.getBase(1) != baseValue {
		t.Errorf("expected tail initial value to be %v, got %v", baseValue, d.getBase(1))
	}
}

func TestGetSetBase(t *testing.T) {
	d := NewDoubleArrayTrie()

	d.setBase(2, 5)

	if d.getBase(2) != 5 {
		t.Errorf("expected base array value at 2 to be %v, got %v", 5, d.getBase(2))
	}
}

func TestGetSetCheck(t *testing.T) {
	d := NewDoubleArrayTrie()

	d.setCheck(2, 5)

	if d.getCheck(2) != 5 {
		t.Errorf("expected check array value at 2 to be %v, got %v", 5, d.getCheck(2))
	}
}

func TestReadTailZeroIndex(t *testing.T) {
	d := NewDoubleArrayTrie()

	if d.ReadTail(1) != "" {
		t.Errorf("expected tail array value at 0 to be %v, got %v", "", d.ReadTail(1))
	}
}

func TestReadTailNonZeroIndex(t *testing.T) {
	d := NewDoubleArrayTrie()

	if d.ReadTail(2) != "" {
		t.Errorf("expected tail array value at 0 to be %v, got %v", "", d.ReadTail(1))
	}
}

func TestReadTailNonZeroTail(t *testing.T) {
	d := NewDoubleArrayTrie()

	d.WriteTail("Hello#", 1)

	if d.ReadTail(1) != "Hello" {
		t.Errorf("expected tail array value at 0 to be %v, got %v", "Hello", d.ReadTail(1))
	}
}

func TestReadTailNonZeroTailMultiple(t *testing.T) {
	d := NewDoubleArrayTrie()

	d.WriteTail("Hello#", 1)
	d.WriteTail("World#", 7)

	if d.ReadTail(7) != "World" {
		t.Errorf("expected tail array value starting at 7 to be %v, got %v", "World", d.ReadTail(7))
	}
}

func TestWriteTailInitial(t *testing.T) {
	d := NewDoubleArrayTrie()

	d.WriteTail("hello#", d.tailPos)

	if d.tail != "hello#" {
		t.Errorf("expected tail array value to be %v, got %v", "hello#", d.tail)
	}
}

func TestWriteTailNoOverlapping(t *testing.T) {
	d := NewDoubleArrayTrie()

	d.WriteTail("hello#", 1)
	d.WriteTail("world#", 7)

	if d.tail != "hello#world#" {
		t.Errorf("expected tail array value to be %v, got %v", "hello#world#", d.tail)
	}
}

func TestWriteTailOverlapping(t *testing.T) {
	d := NewDoubleArrayTrie()

	d.WriteTail("hello#", 1)
	d.WriteTail("world#", 3)

	if d.tail != "held#o#" {
		t.Errorf("expected tail array value to be %v, got %v", "held#o#", d.tail)
	}
}

func TestGetKeyExistsInTrie(t *testing.T) {
	d := NewDoubleArrayTrie()
	d.Add("baby")

	if d.Get("baby") != true {
		t.Errorf("expected search for key %v to be %v, got %v", "baby", true, false)
	}
}

func TestGetKeyDeleteInTrie(t *testing.T) {
	d := NewDoubleArrayTrie()
	d.Add("baby")

	if d.Delete("baby") != true {
		t.Errorf("expected delete for key %v to be %v, got %v", "babe", true, false)
	}

	if d.Get("baby") != false {
		t.Errorf("expected search for key %v to be %v, got %v", "babe", false, true)
	}
}

func TestFindArcsInTrieSimple(t *testing.T) {
	d := NewDoubleArrayTrie()
	d.setCheck(3, 1)
	d.setBase(3, 1)
	d.setCheck(2, 3)
	d.setBase(2, -1)
	d.WriteTail("aby#", 1)

	if d.findArcs(3)[0] != 1 {
		t.Errorf("expected findArcs for pos %v to be %v, got %v", 3, 1, d.findArcs(3)[0])
	}
}

func TestFindArcsInTrieMultiple(t *testing.T) {
	d := NewDoubleArrayTrie()

	d.setBase(1, 1)
	d.setBase(2, 1)
	d.setBase(3, -1)
	d.setBase(4, -12)
	d.setBase(10, -9)

	d.setCheck(1, 3)
	d.setCheck(2, 1)
	d.setCheck(3, 2)
	d.setCheck(4, 2)
	d.setCheck(10, 1)

	if d.findArcs(2)[0] != 2 {
		t.Errorf("expected findArcs first value for pos %v to be %v, got %v", 2, 2, d.findArcs(2)[0])
	}

	if d.findArcs(2)[1] != 3 {
		t.Errorf("expected findArcs second value for pos %v to be %v, got %v", 2, 3, d.findArcs(2)[1])
	}
}

func TestXCheckInTrie(t *testing.T) {
	d := NewDoubleArrayTrie()

	d.setCheck(1, 3)
	d.setCheck(2, 1)
	d.setCheck(3, 2)
	d.setCheck(4, 2)
	d.setCheck(10, 1)

	if d.xCheck([]int{1}) != 4 {
		t.Errorf("expected xCheck for list %v to be %v, got %v", []int{1}, 4, d.xCheck([]int{1}))
	}
}

func TestXCheckInTrieMultipleNoMatch(t *testing.T) {
	d := NewDoubleArrayTrie()

	d.setCheck(1, 3)
	d.setCheck(2, 1)
	d.setCheck(10, 1)

	if d.xCheck([]int{2, 10}) != 1 {
		t.Errorf("expected xCheck for list %v to be %v, got %v", []int{2, 10}, 1, d.xCheck([]int{2, 10}))
	}
}

func TestXCheckInTrieMultipleWithMatch(t *testing.T) {
	d := NewDoubleArrayTrie()

	d.setCheck(1, 3)
	d.setCheck(2, 1)
	d.setCheck(3, 2)
	d.setCheck(4, 2)
	d.setCheck(10, 1)

	if d.xCheck([]int{2, 10}) != 3 {
		t.Errorf("expected xCheck for list %v to be %v, got %v", []int{2, 10}, 3, d.xCheck([]int{2, 10}))
	}
}

func TestXAddInTrieEmpty(t *testing.T) {
	d := NewDoubleArrayTrie()

	d.Add("bachelor")
	d.Add("jar")

	if d.tail != "achelor#ar#" {
		t.Errorf("expected tail to be %v, got %v", "achelor#ar#", d.tail)
	}

	if d.tailPos != 12 {
		t.Errorf("expected tailPos to be %v, got %v", 11, d.tailPos)
	}

	if d.getBase(3) != -1 {
		t.Errorf("expected getBase for pos %v to be %v, got %v", 3, -1, d.getBase(3))
	}

	if d.getCheck(3) != 1 {
		t.Errorf("expected getCheck for pos %v to be %v, got %v", 3, 1, d.getCheck(3))
	}
}


func TestXAddInTrieWithNoCommonPrefix(t *testing.T) {
	d := NewDoubleArrayTrie()

	d.Add("bachelor")
	d.Add("jar")

	if d.Get("bachelor") != true {
		t.Errorf("expected Get for % to be %v, got %v", "bachelor", true, d.Get("bachelor"))
	}

	if d.Get("jar") != true {
		t.Errorf("expected Get for % to be %v, got %v", "jar", true, d.Get("jar"))
	}
}

func TestXAddInTrieWithCommonPrefix(t *testing.T) {
	d := NewDoubleArrayTrie()

	d.Add("bachelor")
	d.Add("jar")
	d.Add("badge")

	if d.Get("bachelor") != true {
		t.Errorf("expected Get for % to be %v, got %v", "bachelor", true, d.Get("bachelor"))
	}

	if d.Get("jar") != true {
		t.Errorf("expected Get for % to be %v, got %v", "jar", true, d.Get("jar"))
	}

	if d.Get("badge") != true {
		t.Errorf("expected Get for % to be %v, got %v", "badge", true, d.Get("badge"))
	}
}

func BenchmarkDoubleArrayTrieGetSimpleStringKey(b *testing.B) {
	d := NewDoubleArrayTrie()

	words := [...]string{"hellohasdhwd ed  qqdwd", "baby", "are", "you", "today"}
	for _, word := range words {
		d.Add(word)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		d.Get(words[i%len(words)])
	}
}
