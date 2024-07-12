package main

import (
	"fmt"
	"go-test/algorithm"
	// Replace "your-package-path" with the actual package path
)

func main() {
	permuteOut := algorithm.Permutation("abc")
	findOddOut := algorithm.FindOdd([]int{1, 1, 1, 2, 2, 3, 3})
	smiles := algorithm.CountSmileFace([]string{":D"})

	fmt.Println(permuteOut)
	fmt.Println(findOddOut)
	fmt.Println(smiles)
}
