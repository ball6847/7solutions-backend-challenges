package challenge1

import (
	"7solution/util"
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

// analyzeMaxPathSum takes dataset and analyze Maximum Path Sum in a Triangle using bottom-up dynamic programming
func analyzeMaxPathSum(data [][]int) int {
	total := len(data)

	// initialize accumulated values using values from the last row
	acc := data[total-1]

	// loop from second-last row to get highest sum for each values against the previously accumulated values
	for i := total - 1; i >= 1; i-- {
		curr := []int{}
		for n, val := range data[i-1] {
			// pick left or right value from connected accumulated node using max() function
			// add it to active node and append it to temporary accumulated values for current row
			curr = append(curr, val+util.Max(acc[n], acc[n+1]))
		}
		// promote accumulated values for the next iteration
		acc = curr
	}

	// if dataset is correct, we should end up with a single item slice, and that's the answer to the problem
	return acc[0]
}

// Handler for cobra command, the handler read dataset and analyze it
func Handler(cmd *cobra.Command, args []string) {
	// read input from json e
	data, err := readInputFile()
	if err != nil {
		log.Fatalln("Error reading input", err)
	}

	result := analyzeMaxPathSum(data)

	fmt.Println("The answer for challenge #1 is: ", result)
}
