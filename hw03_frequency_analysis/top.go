package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type WordInfo struct {
	Word  string
	Count int
}

func Top10(text string) []string {
	countMap := make(map[string]int)
	textSlice := strings.Fields(text)
	for i := 0; i < len(textSlice); i++ {
		countMap[textSlice[i]]++
	}

	i := 0
	wordsSlice := make([]WordInfo, len(countMap))
	for k, v := range countMap {
		wordsSlice[i] = WordInfo{k, v}
		i++
	}

	sort.Slice(wordsSlice, func(i, j int) bool {
		if wordsSlice[i].Count == wordsSlice[j].Count {
			return strings.Compare(wordsSlice[i].Word, wordsSlice[j].Word) < 0
		}
		return wordsSlice[i].Count > wordsSlice[j].Count
	})

	result := make([]string, len(wordsSlice))
	i = 0
	for _, v := range wordsSlice {
		result[i] = v.Word
		i++
	}

	if len(wordsSlice) > 10 {
		return result[:10]
	}
	return result
}
