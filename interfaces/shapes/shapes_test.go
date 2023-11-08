package shapes

import (
	"math"
	"reflect"
	"testing"
)

func TestArea(t *testing.T) {
	tests := []struct {
		testName string
		input    Shape
		want     float64
	}{
		{testName: "baseCircle", input: Circle{R: 10}, want: math.Pi * 10 * 10},
		{testName: "baseRectangle", input: Rectangle{Width: 10, Height: 20}, want: 10 * 20},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {

			got := tc.input.Area()

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("expected: %#v, got: %#v", tc.want, got)
			}
		})
	}
}
