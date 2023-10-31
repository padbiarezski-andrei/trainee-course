package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	tests := map[string]struct {
		want string
	}{
		"base": {want: "Hello, world!"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Hello()
			if tc.want != got {
				t.Errorf("expected: %#v, got: %#v", tc.want, got)
			}
		})
	}
}
