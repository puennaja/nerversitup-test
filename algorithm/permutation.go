package algorithm

import (
	"slices"
	"strings"
)

func permute(strs []string, filterStrs []string, out *[]string) {
	if len(filterStrs) == 0 {
		join := strings.Join(strs, "")
		if !slices.Contains(*out, join) {
			*out = append(*out, join)
		}
	}

	for i, v := range filterStrs {
		nextStrs := append([]string{}, strs...)
		nextStrs = append(nextStrs, v)
		nextFilterStrs := append([]string{}, filterStrs[:i]...)
		nextFilterStrs = append(nextFilterStrs, filterStrs[i+1:]...)
		permute(nextStrs, nextFilterStrs, out)
	}
}

func Permutation(str string) []string {
	data := strings.Split(str, "")
	out := make([]string, 0)
	permute([]string{}, data, &out)
	return out
}
