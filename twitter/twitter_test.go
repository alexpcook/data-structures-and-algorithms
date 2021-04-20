package twitter

import "testing"

func TestGetFirstAndLastTweets(t *testing.T) {
	tests := []struct {
		tweets      []Tweet
		first, last Tweet
	}{
		{[]Tweet{{"abc", 2010}, {"def", 2016}, {"ghi", 2020}}, Tweet{"abc", 2010}, Tweet{"ghi", 2020}},
		{[]Tweet{{"abc", 2010}}, Tweet{"abc", 2010}, Tweet{"abc", 2010}},
	}
	for _, test := range tests {
		if first, last := GetFirstAndLastTweets(test.tweets); first != test.first || last != test.last {
			t.Fatalf("want %v %v, got %v %v", test.first, test.last, first, last)
		}
	}
}

func BenchmarkGetFirstAndLastTweets(b *testing.B) {
	test := []Tweet{{"abc", 2010}, {"def", 2016}, {"ghi", 2020}}
	for i := 0; i < b.N; i++ {
		GetFirstAndLastTweets(test)
	}
}

func TestGetTimeDiffBetweenFirstAndLastTweets1(t *testing.T) {
	tests := getTimeDiffTests()
	for _, test := range tests {
		if got := GetTimeDiffBetweenFirstAndLastTweets1(test.input); got != test.want {
			t.Fatalf("want %v, got %v", test.want, got)
		}
	}
}

func BenchmarkGetTimeDiffBetweenFirstAndLastTweets1(b *testing.B) {
	test := getTimeDiffTests()[0]
	for i := 0; i < b.N; i++ {
		GetTimeDiffBetweenFirstAndLastTweets1(test.input)
	}
}

func TestGetTimeDiffBetweenFirstAndLastTweets2(t *testing.T) {
	tests := getTimeDiffTests()
	for _, test := range tests {
		if got := GetTimeDiffBetweenFirstAndLastTweets2(test.input); got != test.want {
			t.Fatalf("want %v, got %v", test.want, got)
		}
	}
}

func BenchmarkGetTimeDiffBetweenFirstAndLastTweets2(b *testing.B) {
	test := getTimeDiffTests()[0]
	for i := 0; i < b.N; i++ {
		GetTimeDiffBetweenFirstAndLastTweets2(test.input)
	}
}

type timeDiffTest struct {
	input []Tweet
	want  int
}

func getTimeDiffTests() []timeDiffTest {
	return []timeDiffTest{
		{[]Tweet{{"abc", 2010}, {"def", 2016}, {"ghi", 2020}}, 10},
		{[]Tweet{{"abc", 2010}}, 0},
	}
}
