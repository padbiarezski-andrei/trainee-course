package interfaces

import (
	"bytes"
	"reflect"
	"testing"
)

func TestInterfacesLimitReader(t *testing.T) {
	baseTest := bytes.NewBufferString(`abcdefghijklmnopqrstuvwxyz`)
	tests := []struct {
		testName string
		input    struct {
			count  int64
			reader bytes.Buffer
		}
		want []byte
	}{
		{
			testName: `base`,
			input: struct {
				count  int64
				reader bytes.Buffer
			}{
				count:  10,
				reader: *baseTest,
			},
			want: []byte(`abcdefghij`),
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {

			l := LimitReader(&tc.input.reader, tc.input.count)
			gotBuf := make([]byte, tc.input.count)

			_, err := l.Read(gotBuf)

			if err != nil {
				t.Fatalf("unexpected error: err %s, %#v, got: %#v", err, tc.want, gotBuf)
			}

			if !reflect.DeepEqual(gotBuf, tc.want) {
				t.Fatalf("expected: %#v, got: %#v", tc.want, gotBuf)
			}
		})
	}
}
