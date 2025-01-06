package service

import (
	"regexp"
	"strings"
)

type WordCounter struct{}

// generateWordGroups
func (wc *WordCounter) CountAllWord(text string) map[string]int {
	// convert to lower case for stable result
	cleaned := strings.ToLower(text)

	// remove all punctuation
	re := regexp.MustCompile(`[^\w\s-]`)
	cleaned = re.ReplaceAllString(cleaned, "")

	// generate slice of words by splitting by space
	words := strings.Fields(cleaned)

	// create a map with word as key and count as value
	sum := make(map[string]int)
	for i := 0; i <= len(words)-1; i++ {
		word := words[i]
		sum[word]++
	}

	return sum
}
