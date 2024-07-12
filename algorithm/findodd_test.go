package algorithm_test

import (
	"go-test/algorithm"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOdd(t *testing.T) {
	type caseTest struct {
		name     string
		args     []int
		expected int
	}

	cases := []caseTest{
		{
			name:     "basic case: 1",
			args:     []int{7},
			expected: 7,
		},
		{
			name:     "basic case: 2",
			args:     []int{0, 1, 0, 1, 0},
			expected: 0,
		},
		{
			name:     "basic case: 1",
			args:     []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1},
			expected: 4,
		},
		{
			name:     "basic: no odd find",
			args:     []int{2, 2},
			expected: -1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := algorithm.FindOdd(tc.args)
			assert.Equal(t, tc.expected, actual)
		})
	}

}
