package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		testName string
		want     string
	}{
		{testName: "base", want: "Hello, world!"},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			got := Hello()
			if tc.want != got {
				t.Errorf("expected: %#v, got: %#v", tc.want, got)
			}
		})
	}
}
