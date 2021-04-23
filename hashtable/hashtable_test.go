package hashtable

import "testing"

func TestBasicHashTableOperations(t *testing.T) {
	BasicHashTableOperations()
}

func TestFirstRecurringCharacter(t *testing.T) {
	tests := []struct {
		data    []int
		want    int
		isError bool
	}{
		{[]int{2, 5, 1, 2, 5, 3, 6}, 2, false},
		{[]int{2, 1, 1, 2, 3, 5}, 1, false},
		{[]int{2, 3, 4, 5}, 0, true},
	}

	for _, test := range tests {
		got, err := FirstRecurringCharacter(test.data)

		if test.isError {
			if err == nil {
				t.Fatal("want error, got nil")
			}
			return
		}

		if got != test.want {
			t.Fatalf("want %v, got %v", test.want, got)
		}
	}
}
