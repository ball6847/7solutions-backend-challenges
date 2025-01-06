package service_test

import (
	"7solution/api/service"
	"reflect"
	"testing"
)

type WordCounterTestCase struct {
	name     string
	text     string
	expected map[string]int
}

func TestWordCounter(t *testing.T) {
	testCases := []WordCounterTestCase{
		{
			name: "CountAllWord",
			text: "HEllo world, hello world.",
			expected: map[string]int{
				"hello": 2,
				"world": 2,
			},
		},
	}

	wc := service.WordCounter{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := wc.CountAllWord(tc.text)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v but got %v", tc.expected, result)
			}
		})
	}
}
