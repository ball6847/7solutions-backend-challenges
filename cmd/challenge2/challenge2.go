package challenge2

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type DecodeResult struct {
	solutions []string
	answer    string
}

// isValid check if input and output are align with our constraints
func isValid(in string, out string, partial bool) bool {
	input := ""

	// strip input to match with output length+1 to allow partially checking if needed
	if partial {
		input = in[:len(out)-1]
	} else {
		input = in
	}

	// the length of decoded number must be longer than the input exactly one char
	if len(out)-len(input) != 1 {
		return false
	}

	for i, char := range input {
		// try to parse left and right number as int so we can do numerical comparison
		left, err := strconv.Atoi(string(out[i]))
		if err != nil {
			return false
		}

		right, err := strconv.Atoi(string(out[i+1]))
		if err != nil {
			return false
		}

		// "L" means number on the left is greater than number on the right
		// "R" means number on the right is greater than number on the left
		// "=" means number on the left is equal to number on the right
		// any unexpected char will immediately make string invalid
		if char == 'L' && left > right {
			continue
		} else if char == 'R' && left < right {
			continue
		} else if char == '=' && left == right {
			continue
		} else {
			return false
		}
	}

	// all constraints pass
	return true
}

// backtrack running to find all the possible solutions for a given input based on our constraints
func backtrack(in string, out string, solutions *[]string) string {
	if len(out)-len(in) == 1 {
		if isValid(in, out, false) {
			*solutions = append(*solutions, out)
		}
		return ""
	}

	for i := 0; i < 10; i++ {
		candidate := out + strconv.Itoa(i)
		if isValid(in, candidate, true) {
			res := backtrack(in, candidate, solutions)
			if res != "" {
				return res
			}
		}
	}

	return ""
}

// getLowestSumSolution takes all solutions to find only one with the lowest sum value
func getLowestSumSolution(solutions []string) (string, error) {
	lowest := 0
	solution := ""
	for _, s := range solutions {
		sum, err := numericStringSum(s)
		if err != nil {
			return "", fmt.Errorf("failed to calculate sum for numeric string %q", s)
		}
		if lowest == 0 || sum < lowest {
			lowest = sum
			solution = s
		}
	}
	return solution, nil
}

// numericStringSum parses numeric string as digits and add them together (sum)
func numericStringSum(num string) (int, error) {
	total := 0
	for _, s := range num {
		n, err := strconv.Atoi(string(s))
		if err != nil {
			return 0, err
		}
		total += n
	}
	return total, nil
}

func dumpSolution(solutions *[]string) error {
	// Join the slice into a single string with newlines
	content := strings.Join(*solutions, "\n")

	// Write the string to a file
	err := os.WriteFile("solutions.txt", []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

// Decode run input against our backtrack algorithm
// returns all possible solutions from backtrack along with the best answer
func Decode(in string) (DecodeResult, error) {
	solutions := []string{}
	backtrack(in, "", &solutions)

	if len(solutions) == 0 {
		return DecodeResult{}, fmt.Errorf("cannot found solution for %q", in)
	}

	answer, err := getLowestSumSolution(solutions)
	if err != nil {
		return DecodeResult{}, err
	}

	return DecodeResult{
		solutions: solutions,
		answer:    answer,
	}, nil
}

// Handler for cobra command
func Handler(cmd *cobra.Command, args []string) {
	in := ""

	// flags
	dump, _ := cmd.Flags().GetBool("dump-solutions")

	// prompt user to enter input
	fmt.Print("Please enter your encoded input:")
	fmt.Scanln(&in)

	// perform decoding
	out, err := Decode(in)
	if err != nil {
		log.Fatalf("failed to decode %q : %v\n", in, err)
	}

	// answer
	fmt.Printf("Total solution: %d\n", len(out.solutions))
	fmt.Printf("Best solution: %v\n", out.answer)

	// dump solutions if needed
	if dump {
		err := dumpSolution(&out.solutions)
		if err != nil {
			fmt.Println("Failed to dump solutions to solutions.txt: ", err)
		} else {
			fmt.Println("All solutions has been dumped to solutions.txt")
		}
	}
}
