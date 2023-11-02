package palindrome

import (
	"testing"
)

func BenchmarkPalindrome1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Palindrome1("test")
	}
}

func BenchmarkPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Palindrome2("test")
	}
}

func BenchmarkPalindrome3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Palindrome3("test")
	}
}

var palindromeTests = []struct {
	testName string
	input    string
	want     bool
}{
	{testName: "empty", input: "", want: true},     // ???
	{testName: "one char", input: "!", want: true}, // ???
	{testName: "true test 1", input: "22\\2\\22", want: true},
	{testName: "true test 2", input: "Mr. Owl ate my metal worm", want: true},
	{testName: "false test 1", input: "youtube", want: false},
}

func TestPalindrome1(t *testing.T) {
	// t.Parallel()

	for _, tc := range palindromeTests {
		t.Run(tc.testName, func(t *testing.T) {
			got := Palindrome1(tc.input)
			if tc.want != got {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}

func TestPalindrome2(t *testing.T) {
	// t.Parallel()

	for _, tc := range palindromeTests {
		t.Run(tc.testName, func(t *testing.T) {
			got := Palindrome2(tc.input)
			if tc.want != got {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}

func TestPalindrome3(t *testing.T) {
	// t.Parallel()

	for _, tc := range palindromeTests {
		t.Run(tc.testName, func(t *testing.T) {
			got := Palindrome3(tc.input)
			if tc.want != got {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
