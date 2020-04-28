package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the [delete] command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
