package algorithm_test

import (
	"go-test/algorithm"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutaion(t *testing.T) {
	type caseTest struct {
		name     string
		args     string
		expected []string
	}

	cases := []caseTest{
		{
			name:     "basic case: 1",
			args:     "a",
			expected: []string{"a"},
		},
		{
			name:     "basic case: 2",
			args:     "ab",
			expected: []string{"ab", "ba"},
		},
		{
			name:     "basic case: 3",
			args:     "aabb",
			expected: []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := algorithm.Permutation(tc.args)
			assert.Equal(t, tc.expected, actual)
		})
	}

}
