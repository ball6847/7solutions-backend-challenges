package challenge3

import (
	"7solution/api"

	"github.com/spf13/cobra"
)

func Handler(cmd *cobra.Command, args []string) {
	api.Serve()
}
