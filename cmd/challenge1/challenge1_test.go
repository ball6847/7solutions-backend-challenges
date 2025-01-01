package challenge1

import (
	"strings"
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

	output, _ := analyzeHighestValuePath(input)
	if output != expected {
		t.Fatalf(`output value should be %d, %d given`, expected, output)
	}
}

func TestAnalyzeHightValuePathDataError1(t *testing.T) {
	input := [][]int{
		{59},
		{73},
	}
	expectedErr := "each row should has one more item added"

	_, err := analyzeHighestValuePath(input)
	if err == nil {
		t.Fatalf("expected an error but got nil")
	}

	if !strings.Contains(err.Error(), expectedErr) {
		t.Fatalf("expected error message to contain %q, but got %q", expectedErr, err.Error())
	}
}

func TestAnalyzeHightValuePathDataError2(t *testing.T) {
	input := [][]int{
		{59},
		{73, 73},
	}
	expectedErr := "duplicated values block our way to proceed next row"

	_, err := analyzeHighestValuePath(input)
	if err == nil {
		t.Fatalf("expected an error but got nil")
	}

	if !strings.Contains(err.Error(), expectedErr) {
		t.Fatalf("expected error message to contain %q, but got %q", expectedErr, err.Error())
	}
}
