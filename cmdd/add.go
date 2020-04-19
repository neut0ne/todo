package cmdd

import (
  "fmt"
  "github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
  Use:   "add",
  Short: "Adds a task to your task list.",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Add called")
  },
}

// init will run before main. Good for setting up stuff before start.
// Runs on package level vars and funcs.
func init() {
  RootCmd.AddCommand(addCmd)
}
