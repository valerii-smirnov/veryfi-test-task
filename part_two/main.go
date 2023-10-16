package main

import (
	"fmt"
	"github.com/valerii-smirnov/veryfi-test-task/part_two/anagrams"
)

func main() {
	in := []string{"affx", "a", "ab", "ba", "nnx", "xnn", "cde", "edc", "dce", "xffa"}
	res := anagrams.GroupAnagrams(in)

	fmt.Println(res)
}
