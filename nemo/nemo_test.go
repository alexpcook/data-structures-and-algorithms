package nemo

import (
	"testing"
)

func getNemoSlice(length int) []string {
	slc := make([]string, length)
	for i := range slc {
		if i%100 == 0 {
			slc[i] = "nemo"
		} else {
			slc[i] = "dory"
		}
	}
	return slc
}

func getFindNemoTests() [][]string {
	return [][]string{
		{"marlin", "dory", "nemo", "crush", "squirt"},
		getNemoSlice(1e3),
		getNemoSlice(1e4),
		getNemoSlice(1e5),
	}
}

func TestFindNemo(t *testing.T) {
	tests := getFindNemoTests()
	for _, test := range tests {
		FindNemo(test)
	}
}

func BenchmarkFindNemo(b *testing.B) {
	tests := getFindNemoTests()
	test := tests[1]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindNemo(test)
	}
}
