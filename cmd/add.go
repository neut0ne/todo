package cmd

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
// Runs on package level vars and funcs. Might be less good for ... (-13.26)
func init() {
  // And this worked without importing cmdd because we're already in that
  // package, so commands in files in the package is already available!
  // Lightbulb moment :P
  RootCmd.AddCommand(addCmd)
}
