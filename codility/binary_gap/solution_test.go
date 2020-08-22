package binary_gap

import (
	"testing"
)

func TestBinaryGap(t *testing.T) {
	testCases := map[int]int{
		1041: 5,
		529:  4,
		32:   0,
		20:   1,
		9:    2,
	}

	for input, expectation := range testCases {
		output := BinaryGap(input)
		if expectation != output {
			t.Errorf("Expected binary gap for %d is %d, but %d received instead", input, expectation, output)
		}
	}
}
