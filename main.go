package main

import (
	"7solution/cmd/challenge1"
	"log"

	"github.com/spf13/cobra"
)

func main() {
	app := &cobra.Command{Use: "7solutions"}

	// register commands
	app.AddCommand(
		&cobra.Command{
			Use:   "challenge1",
			Short: "Run and see the result for challenge1",
			Run:   challenge1.Handler,
		},
	)

	if err := app.Execute(); err != nil {
		log.Fatalln("Error starting cli application", err)
	}
}
