package algorithm

import (
	"slices"
	"strings"
)

var (
	eyes   = []string{":", ";"}
	nose   = []string{"-", "~"}
	smiles = []string{")", "D"}
)

func CountSmileFace(faces []string) int {
	countSmile := 0
	for _, face := range faces {
		sFace := strings.Split(face, "")
		if len(sFace) == 2 {
			if slices.Contains(eyes, sFace[0]) && slices.Contains(smiles, sFace[1]) {
				countSmile++
			}
		} else if len(sFace) == 3 {
			if slices.Contains(eyes, sFace[0]) && slices.Contains(nose, sFace[1]) && slices.Contains(smiles, sFace[2]) {
				countSmile++
			}
		} else {
			continue
		}
	}
	return countSmile
}
