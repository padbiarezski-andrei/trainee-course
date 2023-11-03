package duplicates

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		testName string
		input    []string
		want     []string
	}{
		{testName: "no dup", input: []string{"India", "Canada", "Japan", "Germany", "Italy"}, want: []string{"India", "Canada", "Japan", "Germany", "Italy"}},
		{testName: "base", input: []string{"India", "India", "Canada", "Japan", "Canada", "Japan", "Germany", "Italy", "Italy"}, want: []string{"India", "Canada", "Japan", "Germany", "Italy"}},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			got := removeDuplicates(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}

}
