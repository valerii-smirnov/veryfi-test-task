package anagrams

import "strings"

// GroupAnagrams find anagrams in the provided strings slice and return them grouped.
// The function result is not sorted alphabetically for groups and values within groups.
// This is because the algorithm implementation is based on a hash table,
// which does not guarantee the order of elements in Go.
// I've omitted the result sorting to achieve the optimal algorithm complexity of O(N*M).
// I could add sorting to the result of the function, but this would increase the algorithm's complexity to O(N * M log M).
// In this solution, I wanted to implement the most efficient algorithm possible.
func GroupAnagrams(in []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, str := range in {
		str = strings.ToLower(str)

		count := [26]int{}
		for _, ch := range str {
			count[ch-'a']++
		}
		groups[count] = append(groups[count], str)
	}

	ret := make([][]string, 0, len(groups))

	for _, words := range groups {
		ret = append(ret, words)
	}

	return ret
}
