package hashtable

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCustomHashTable(t *testing.T) {
	ht := NewHashTable(2)
	if len(ht.data) != 2 || ht.size != 2 {
		t.Fatalf("expected size of hash table to be 2, got %v, %v", len(ht.data), ht.size)
	}
	if keys := ht.Keys(); !reflect.DeepEqual(keys, []string{}) {
		t.Fatalf("expected empty key set, got %v", keys)
	}
	if values := ht.Values(); !reflect.DeepEqual(values, []string{}) {
		t.Fatalf("expected empty value set, got %v", values)
	}

	fruit := "apple"
	color := "red"
	ht.Set(fruit, color)
	if got := ht.Get(fruit); got != color {
		t.Fatalf("want color %v, got %v", color, got)
	}
	if keys := ht.Keys(); !reflect.DeepEqual(keys, []string{"apple"}) {
		t.Fatalf("expected empty key set, got %v", keys)
	}
	if values := ht.Values(); !reflect.DeepEqual(values, []string{"red"}) {
		t.Fatalf("expected empty value set, got %v", values)
	}

	color = "green"
	ht.Set(fruit, color)
	if got := ht.Get(fruit); got != color {
		t.Fatalf("want color %v, got %v", color, got)
	}
	if keys := ht.Keys(); !reflect.DeepEqual(keys, []string{"apple"}) {
		t.Fatalf("expected empty key set, got %v", keys)
	}
	if values := ht.Values(); !reflect.DeepEqual(values, []string{"green"}) {
		t.Fatalf("expected empty value set, got %v", values)
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
	ht.Delete("apple")
	ht.Delete("peach")
	if keys := ht.Keys(); !reflect.DeepEqual(keys, []string{"banana"}) {
		t.Fatalf("expected empty key set, got %v", keys)
	}
	if values := ht.Values(); !reflect.DeepEqual(values, []string{"yellow"}) {
		t.Fatalf("expected empty value set, got %v", values)
	}

	fmt.Println(ht)
}
