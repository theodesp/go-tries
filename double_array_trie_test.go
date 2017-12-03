package go_tries

import (
	"testing"
)

func TestInitTail(t *testing.T) {
	d := NewDoubleArrayTrie()

	if d.tail != "#" {
		t.Errorf("expected tail initial value to be %v, got %v", "#", d.tail)
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

	if d.ReadTail(0) != "" {
		t.Errorf("expected tail array value at 0 to be %v, got %v", "", d.ReadTail(0))
	}
}

func TestWriteTailInitial(t *testing.T) {
	d := NewDoubleArrayTrie()

	d.WriteTail("hello#", 0)

	if d.tail != "hello#" {
		t.Errorf("expected tail array value to be %v, got %v", "hello#", d.tail)
	}
}

func TestWriteTailNoOverlapping(t *testing.T) {
	d := NewDoubleArrayTrie()

	d.WriteTail("hello#", 0)
	d.WriteTail("world#", 6)

	if d.tail != "hello#world#" {
		t.Errorf("expected tail array value to be %v, got %v", "hello#world#", d.tail)
	}
}

func TestWriteTailOverlapping(t *testing.T) {
	d := NewDoubleArrayTrie()

	d.WriteTail("hello#", 0)
	d.WriteTail("world#", 3)

	if d.tail != "helworld#" {
		t.Errorf("expected tail array value to be %v, got %v", "helworld#", d.tail)
	}
}

func TestGetKeyExistsInTrie(t *testing.T) {
	d := NewDoubleArrayTrie()
	d.setCheck(3, 1)
	d.setBase(3, 1)
	d.setCheck(2, 3)
	d.setBase(2, -1)
	d.WriteTail("aby#", 0)

	if d.Get("baby") != true {
		t.Errorf("expected search for key %v to be %v, got %v", "babe", true, false)
	}
}

func BenchmarkDoubleArrayTrieGetStringKey(b *testing.B) {
	d := NewDoubleArrayTrie()

	words := [...]string{
		"hello", "baby", "how", "are", "you", "babe"}

	d.setCheck(3, 1)
	d.setBase(3, 1)
	d.setCheck(2, 3)
	d.setBase(2, -1)
	d.WriteTail("aby#", 0)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		d.Get(words[i%len(words)])
	}
}
