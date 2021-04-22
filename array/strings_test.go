package array

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		data, want string
	}{
		{"hello", "olleh"},
		{"alex", "xela"},
		{"sophia", "aihpos"},
		{"βЩѥ", "ѥЩβ"},
		{"", ""},
	}

	for i, test := range tests {
		if got := ReverseString(test.data); got != test.want {
			t.Fatalf("test %d: want %v, got %v", i, test.want, got)
		}
	}
}
