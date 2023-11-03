package reverse

import (
	"reflect"
	"testing"
)

func BenchmarkReverseWordsWithoutSTD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseWordsWithoutSTD("qwert asdfg zxcvb   youtube")
	}
}

func BenchmarkReverseWordsWIthStd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseWordsWithSTD("qwert asdfg zxcvb   youtube")
	}
}

var reverseTests = []struct {
	testName string
	input    []rune
	want     []rune
}{
	{testName: "empty", input: []rune(""), want: []rune("")},
	{testName: "one char", input: []rune("a"), want: []rune("a")},
	{testName: "base", input: []rune("youtube"), want: []rune("ebutuoy")},
}

func TestReverseInPlace(t *testing.T) {
	for _, tc := range reverseTests {
		t.Run(tc.testName, func(t *testing.T) {
			reverseInPlace(tc.input)
			if !reflect.DeepEqual(tc.input, tc.want) {
				t.Fatalf("expected: %#v, got: %#v", tc.want, tc.input)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	for _, tc := range reverseTests {
		t.Run(tc.testName, func(t *testing.T) {
			got := reverse(string(tc.input))
			if !reflect.DeepEqual(tc.input, tc.want) {
				t.Fatalf("expected: %#v, got: %#v", tc.want, got)
			}
		})
	}
}

var reverseWordsTests = []struct {
	testName string
	input    string
	want     string
}{
	{testName: "empty", input: "", want: ""},
	{testName: "spaces", input: "     ", want: "     "},
	{testName: "日本語", input: "日本語", want: "語本日"},
	{testName: "one word", input: "qwert", want: "trewq"},
	{testName: "base", input: "qwert asdfg zxcvb   youtube", want: "trewq gfdsa bvcxz   ebutuoy"},
}

func TestReverseWords1(t *testing.T) {
	for _, tc := range reverseWordsTests {
		t.Run(tc.testName, func(t *testing.T) {
			got := reverseWordsWithoutSTD(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("expected: %#v, got: %#v", tc.want, got)
			}
		})
	}
}

func TestReverseWords2(t *testing.T) {
	for _, tc := range reverseWordsTests {
		t.Run(tc.testName, func(t *testing.T) {
			got := reverseWordsWithSTD(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("expected: %#v, got: %#v", tc.want, got)
			}
		})
	}
}
