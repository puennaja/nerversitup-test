package algorithm_test

import (
	"go-test/algorithm"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSmileFace(t *testing.T) {
	type caseTest struct {
		name     string
		args     []string
		expected int
	}

	cases := []caseTest{
		{
			name:     "besic: no smile face find",
			args:     []string{},
			expected: 0,
		},
		{
			name:     "besic case: 1",
			args:     []string{":=)", ":)", ":-D"},
			expected: 2,
		},
		{
			name:     "besic case: 2",
			args:     []string{";]", ":0", "jnjsc", ":-D"},
			expected: 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := algorithm.CountSmileFace(tc.args)
			assert.Equal(t, tc.expected, actual)
		})
	}

}
