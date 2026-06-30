package main

import (
	"fmt"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	// Criamos um mapa onde a chave será a string ordenada e o valor será um slice de strings que são anagramas da chave
	anagramMap := make(map[string][]string)

	for _, s := range strs {
		// Geramos a chave ordenada da string
		key := makeSortedKey(s)
		// Adicionamos a string original ao slice da chave
		anagramMap[key] = append(anagramMap[key], s)
	}

	// Criamos um slice de slices para armazenar o resultado
	result := [][]string{}
	// Iteramos sobre o mapa e adicionamos cada slice de strings ao resultado
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}

func makeSortedKey(s string) string {
	b := []byte(s)

	slices.Sort(b)

	return string(b)
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groups := groupAnagrams(strs)
	fmt.Println(groups)
}
