package challenge2

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type DecodeResult struct {
	solutions []string
	optimal   string
}

// isValid check if input and output are align with our constraints
func isValid(in string, out string, partial bool) bool {
	input := ""

	// strip input to match with output length-1 to allow partially checking if needed
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
		left := int(out[i] - '0')    // Convert char to int
		right := int(out[i+1] - '0') // Convert next char to int

		switch char {
		case 'L':
			if left <= right {
				return false
			}
		case 'R':
			if left >= right {
				return false
			}
		case '=':
			if left != right {
				return false
			}
		default:
			return false // Invalid character in input
		}
	}

	// all constraints pass
	return true
}

// backtrack running to find all the possible solutions for a given input based on our constraints
func backtrack(in, out string, solutions []string) []string {
	if len(out)-len(in) == 1 {
		if isValid(in, out, false) {
			return append(solutions, out)
		}
		return solutions
	}

	for i := 0; i < 10; i++ {
		candidate := fmt.Sprintf("%s%d", out, i)
		if isValid(in, candidate, true) {
			solutions = backtrack(in, candidate, solutions)
		}
	}

	return solutions
}

// getOptimalSolution takes all solutions to find only one with the lowest sum value
func getOptimalSolution(solutions []string) (string, error) {
	value := 0
	solution := ""
	for _, s := range solutions {
		sum, err := numericStringSum(s)
		if err != nil {
			return "", fmt.Errorf("failed to calculate sum for numeric string %q", s)
		}
		if value == 0 || sum < value {
			value = sum
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

func dumpSolution(solutions []string) error {
	// Join the slice into a single string with newlines
	content := strings.Join(solutions, "\n")

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
	if in == "" {
		return DecodeResult{}, errors.New("no input")
	}

	solutions := backtrack(in, "", []string{})

	if len(solutions) == 0 {
		return DecodeResult{}, fmt.Errorf("cannot found solution for %q", in)
	}

	answer, err := getOptimalSolution(solutions)
	if err != nil {
		return DecodeResult{}, err
	}

	return DecodeResult{
		solutions: solutions,
		optimal:   answer,
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
	fmt.Printf("Best solution: %v\n", out.optimal)

	// dump solutions if needed
	if dump {
		err := dumpSolution(out.solutions)
		if err != nil {
			fmt.Println("Failed to dump solutions to solutions.txt: ", err)
		} else {
			fmt.Println("All solutions has been dumped to solutions.txt")
		}
	}
}
