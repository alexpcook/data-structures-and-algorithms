package hashtable

import (
	"testing"
)

func TestCustomHashTable(t *testing.T) {
	ht := NewHashTable(2)
	if len(ht.data) != 2 || ht.size != 2 {
		t.Fatalf("expected size of hash table to be 2, got %v, %v", len(ht.data), ht.size)
	}

	fruit := "apple"
	color := "red"
	ht.Set(fruit, color)
	if got := ht.Get(fruit); got != color {
		t.Fatalf("want color %v, got %v", color, got)
	}

	color = "green"
	ht.Set(fruit, color)
	if got := ht.Get(fruit); got != color {
		t.Fatalf("want color %v, got %v", color, got)
	}

	fruit = "banana"
	color = "yellow"
	ht.Set(fruit, color)
	if got := ht.Get(fruit); got != color {
		t.Fatalf("want color %v, got %v", color, got)
	}

	fruit = "peach"
	color = "orange"
	ht.Set(fruit, color)
	if got := ht.Get(fruit); got != color {
		t.Fatalf("want color %v, got %v", color, got)
	}

	if got := ht.Get("somenonexistentkey"); got != "" {
		t.Fatalf("want \"\", got %v", got)
	}

	ht.Delete("somenonexistentkey")
}
