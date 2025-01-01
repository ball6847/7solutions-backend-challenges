package challenge1

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// readInputFile read input json file (expected hard.json at current working directory)
// and returns dataset for further analysis
func readInputFile() ([][]int, error) {
	file, err := os.Open("hard.json")
	if err != nil {
		return nil, errors.New("failed to open input file hard.json")
	}
	defer file.Close()

	var data [][]int
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return nil, errors.New("failed to decode input as json file")
	}

	return data, nil
}

// analyzeHighestValuePath takes dataset and analyze each row
// -to get the best path with highest possible value while keeping each node connected
func analyzeHighestValuePath(data [][]int) (int, error) {
	i := 0
	value := 0
	prevLength := 0

	// iterate over each row
	for row, col := range data {
		// validate row len, as we can't work with invalid dataset
		length := len(col)
		if (length - prevLength) != 1 {
			return 0, fmt.Errorf("invalid dataset found on row index %d: each row should has one more item added", row)
		}

		prevLength = length

		// handle first row, as it always contain single value
		if row == 0 {
			value += col[0]
			continue
		}

		// NOTE: the requirement doesn't say anything about conflicted values, we assume to stop the process and return error
		// TODO: need to clarify this with PO
		if col[i] == col[i+1] {
			return 0, fmt.Errorf("invalid dataset found on row index %d: duplicated values block our way to proceed next row", row)
		} else if col[i] > col[i+1] {
			// select the highest value between I and I+1
			value += col[i]
		} else {
			// position changed to next column as value on the right has higher value
			value += col[i+1]
			i += 1
		}
	}

	return value, nil
}

// Handler for cobra command, the handler read dataset and analyze it
func Handler(cmd *cobra.Command, args []string) {
	// read input from json file
	data, err := readInputFile()
	if err != nil {
		log.Fatalln("Error reading input", err)
	}

	result, err := analyzeHighestValuePath(data)
	if err != nil {
		log.Fatalln("Error analyzing input", err)
	}

	fmt.Println("The value for the best path is: ", result)
}
