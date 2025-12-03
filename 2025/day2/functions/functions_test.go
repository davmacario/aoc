package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInvalid2(t *testing.T) {
	testcases := []struct {
		input  int
		output bool
	}{
		{
			input:  999,
			output: true,
		},
		{
			input:  123123,
			output: true,
		},
		{
			input:  1231,
			output: false,
		},
	}

	for _, tc := range testcases {
		res := IsInvalid2(tc.input)
		assert.Equal(t, res, tc.output)
	}
}
