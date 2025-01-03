package main

import (
	"7solution/cmd/challenge1"
	"7solution/cmd/challenge2"
	"7solution/cmd/challenge3"
	"log"

	"github.com/spf13/cobra"
)

func main() {
	app := &cobra.Command{Use: "7solutions"}

	challenge1Cmd := &cobra.Command{
		Use:   "challenge1",
		Short: "Run and see the result for challenge1",
		Run:   challenge1.Handler,
	}

	challenge2Cmd := &cobra.Command{
		Use:   "challenge2",
		Short: "Run and see the result for challenge2",
		Run:   challenge2.Handler,
	}

	challenge2Cmd.Flags().Bool("dump-solutions", false, "Dump all solutions to a solutions.txt file at current directory")

	challenge3Cmd := &cobra.Command{
		Use:   "challenge3",
		Short: "Start api server for challenge3",
		Run:   challenge3.Handler,
	}

	// register commands
	app.AddCommand(
		challenge1Cmd,
		challenge2Cmd,
		challenge3Cmd,
	)

	if err := app.Execute(); err != nil {
		log.Fatalln("Error starting cli application", err)
	}
}
