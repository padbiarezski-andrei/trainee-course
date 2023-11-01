package duplicates

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	// t.Parallel()

	tests := map[string]struct {
		input []string
		want  []string
	}{
		"no dup": {input: []string{"India", "Canada", "Japan", "Germany", "Italy"}, want: []string{"India", "Canada", "Japan", "Germany", "Italy"}},
		"base":   {input: []string{"India", "India", "Canada", "Japan", "Canada", "Japan", "Germany", "Italy", "Italy"}, want: []string{"India", "Canada", "Japan", "Germany", "Italy"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := removeDuplicates(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}

}
