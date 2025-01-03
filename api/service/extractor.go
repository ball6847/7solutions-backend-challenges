package service

import (
	"7solution/utils"
	"regexp"
	"sort"
	"strings"
)

type Extractor struct{}

type Pair struct {
	Phrase string
	Count  int
}

func NewExtractor() *Extractor {
	return &Extractor{}
}

func (e *Extractor) Extract(input string, min int, excludes []string) []Pair {
	text := strings.ToLower(input)
	var words []Pair
	var pairs []Pair
	// look for 3 words, 2 words then single word
	for n := 3; n >= 1; n-- {
		text = cleanup(text)
		tokens := strings.Fields(text)
		grams := generateNGrams(tokens, n)
		if n > 1 {
			pairs = getFilteredPairs(grams, min, excludes)
		} else {
			// do not apply min occurrence on single word
			pairs = getFilteredPairs(grams, 0, excludes)
		}
		words = append(words, pairs...)

		// remove phrase from original text, so it won't be confused on next round
		for _, pair := range pairs {
			text = strings.ReplaceAll(text, pair.Phrase, "")
		}
	}
	return words
}

// cleanup remove unwanted characters for easier extraction
func cleanup(text string) string {
	// Regex to preserve words with hyphens (e.g., t-bone)
	re := regexp.MustCompile(`[^\w\s-]`) // Remove unwanted characters except hyphen
	cleanedText := re.ReplaceAllString(text, "")
	return cleanedText
}

// Generate n-grams from tokenized words
func generateNGrams(tokens []string, n int) map[string]int {
	ngrams := make(map[string]int)
	for i := 0; i <= len(tokens)-n; i++ {
		ngram := strings.Join(tokens[i:i+n], " ")
		ngrams[ngram]++
	}
	return ngrams
}

// Find the most common n-grams
func getFilteredPairs(ngrams map[string]int, count int, exclude []string) []Pair {
	var pairs []Pair
	for k, v := range ngrams {
		if v >= count && !utils.SliceContains(exclude, k) {
			pairs = append(pairs, Pair{k, v})
		}
	}

	// Sort pairs by count (descending)
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Count > pairs[j].Count
	})

	return pairs
}
