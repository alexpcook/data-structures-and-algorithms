package box

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getOneBox() string {
	runes := make([]rune, rand.Intn(10)+10)
	for i := range runes {
		runes[i] = 'a' + rune(rand.Intn(26))
	}
	return string(runes)
}

func getBoxes(length int) []string {
	slc := make([]string, length)
	for i := range slc {
		slc[i] = getOneBox()
	}
	return slc
}

func getTestCases() [][]string {
	return [][]string{
		{"box1", "box2", "box3", "box3"},
		getBoxes(1e3),
		getBoxes(1e4),
		getBoxes(1e5),
	}
}

func TestLogFirstBox(t *testing.T) {
	tests := getTestCases()
	for _, test := range tests {
		LogFirstBox(test)
	}
}

func BenchmarkLogFirstBox(b *testing.B) {
	tests := getTestCases()
	test := tests[1]
	for i := 0; i < b.N; i++ {
		LogFirstBox(test)
	}
}

func TestLogFirstTwoBoxes(t *testing.T) {
	tests := getTestCases()
	for _, test := range tests {
		LogFirstTwoBoxes(test)
	}
}

func BenchmarkLogFirstTwoBoxes(b *testing.B) {
	tests := getTestCases()
	test := tests[1]
	for i := 0; i < b.N; i++ {
		LogFirstTwoBoxes(test)
	}
}

func TestLogAllPairsOfBoxes(t *testing.T) {
	test := getTestCases()[0]
	LogAllPairsOfBoxes(test)
}

func BenchmarkLogAllPairsOfBoxes(b *testing.B) {
	test := getTestCases()[0]
	for i := 0; i < b.N; i++ {
		LogAllPairsOfBoxes(test)
	}
}

func TestLogBoxesRecursive(t *testing.T) {
	tests := []int{1e0, 1e1}
	for _, test := range tests {
		LogBoxesRecursive(test)
	}
}

func BenchmarkLogBoxesRecursive(b *testing.B) {
	test := 10
	for i := 0; i < b.N; i++ {
		LogBoxesRecursive(test)
	}
}
