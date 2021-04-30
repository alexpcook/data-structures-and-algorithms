package algorithms

import (
	"testing"
	"time"
)

func TestAddTo80(t *testing.T) {
	tests := []struct {
		n, want int
	}{
		{5, 85},
		{0, 80},
		{-19, 61},
	}

	for i, test := range tests {
		if got := AddTo80(test.n); got != test.want {
			t.Fatalf("test %d: want %d, got %d", i, test.want, test.n)
		}
	}
}

func TestMemorizedAddTo80(t *testing.T) {
	data := []int{-19, 0, 5}
	for _, d := range data {
		MemorizedAddTo80(d) // No cache hits yet
	}

	memChan := make(chan int)
	for _, d := range data {
		go func() {
			memChan <- MemorizedAddTo80(d) // These should all be cache hits that take less than five seconds
		}()

		timer := time.NewTimer(1 * time.Second)
		select {
		case result := <-memChan:
			if result != d+80 {
				t.Fatalf("want %d, got %d", d+80, result)
			}
		case <-timer.C:
			t.Fatalf("looking up %d took too long", d)
		}
	}
}
