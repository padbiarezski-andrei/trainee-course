package reverse

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := map[string]struct {
		input []rune
		want  []rune
	}{
		"empty":    {input: []rune(""), want: []rune("")},
		"one char": {input: []rune("a"), want: []rune("a")},
		"base":     {input: []rune("youtube"), want: []rune("ebutuoy")},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			reverse(tc.input)
			if !reflect.DeepEqual(tc.input, tc.want) {
				t.Fatalf("expected: %v, got: %v", tc.input, tc.want)
			}
		})
	}
}

var reverseWordsTests = map[string]struct {
	input string
	want  string
}{
	"empty":    {input: "", want: ""},
	"spaces":   {input: "     ", want: "     "},
	"one word": {input: "qwert", want: "trewq"},
	"base":     {input: "qwert asdfg zxcvb   youtube", want: "trewq gfdsa bvcxz   ebutuoy"},
}

func TestReverseWords(t *testing.T) {
	t.Parallel()
	for name, tc := range reverseWordsTests {
		t.Run(name, func(t *testing.T) {
			got := reverseWords1(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}

func TestReverseWords2(t *testing.T) {
	t.Parallel()
	for name, tc := range reverseWordsTests {
		t.Run(name, func(t *testing.T) {
			got := reverseWords2(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
