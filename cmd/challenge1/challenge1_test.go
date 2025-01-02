package challenge1

import (
	"testing"
)

// TestAnalyzeHighestValuePath
func TestAnalyzeHighestValuePath(t *testing.T) {
	input := [][]int{
		{59},
		{73, 41},
		{52, 40, 53},
		{26, 53, 6, 34},
	}
	expected := 237

	output := analyzeMaxPathSum(input)
	if output != expected {
		t.Fatalf(`output value should be %d, %d given`, expected, output)
	}
}
