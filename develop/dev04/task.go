package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	a := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	fmt.Println(SearchAnagrams(a))
}

type L struct {
	word      string
	lastIndex int
}

func SearchAnagrams(words []string) map[string][]string {
	for i, v := range words {
		words[i] = strings.ToLower(v)
	}
	sortedWords := make([]*L, 0, len(words))
	for i, v := range words {
		runeWord := []rune(v)
		sort.Slice(runeWord, func(i, j int) bool {
			return runeWord[i] < runeWord[j]
		})
		sortedWords = append(sortedWords, &L{string(runeWord), i})
	}
	sort.Slice(sortedWords, func(i, j int) bool {
		return sortedWords[i].word < sortedWords[j].word
	})
	minIndex, currWord := sortedWords[0].lastIndex, sortedWords[0].word
	result, anagrams := make(map[string][]string), []string{words[minIndex]}
	for _, v := range sortedWords[1:] {
		if currWord != v.word {
			if len(anagrams) > 1 {
				sort.Strings(anagrams)
				result[words[minIndex]] = anagrams
			}
			anagrams, minIndex = make([]string, 0), math.MaxInt
		} else if minIndex > v.lastIndex {
			minIndex = v.lastIndex
		}
		anagrams, currWord = append(anagrams, words[v.lastIndex]), v.word
	}
	if len(anagrams) > 1 {
		sort.Strings(anagrams)
		result[words[minIndex]] = anagrams
	}
	return result
}
