package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, s := range strs {
		key := makeSortedKey(s)
		anagramMap[key] = append(anagramMap[key], s)
	}

	result := [][]string{}
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}

func makeSortedKey(s string) string {
	b := []byte(s)

	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	return string(b)
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groups := groupAnagrams(strs)
	fmt.Println(groups)
}
