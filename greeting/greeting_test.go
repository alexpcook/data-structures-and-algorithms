package greeting

import "testing"

func getTestCases() []int {
	return []int{1e0, 1e1, 1e2, 1e3}
}

func TestBoo(t *testing.T) {
	tests := getTestCases()
	for _, test := range tests {
		Boo(test)
	}
}

func BenchmarkBoo(b *testing.B) {
	test := getTestCases()[2]
	for i := 0; i < b.N; i++ {
		Boo(test)
	}
}

func TestHi(t *testing.T) {
	tests := getTestCases()
	for _, test := range tests {
		got := Hi(test)
		if len(got) != test {
			t.Fatalf("expected slice of length %d, but got %d", test, len(got))
		}

		for _, v := range got {
			if v != "hi" {
				t.Fatalf("expected slice element with value \"hi\", but got %q", v)
			}
		}
	}
}

func BenchmarkHi(b *testing.B) {
	test := getTestCases()[2]
	for i := 0; i < b.N; i++ {
		Hi(test)
	}
}
