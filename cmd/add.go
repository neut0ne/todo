package cmd

import (
  "fmt"
  "strings"
  "github.com/spf13/cobra"
  "strings"
)

var addCmd = &cobra.Command{
  Use:   "add",
  Short: "Adds a task to your task list.",
  Run: func(cmd *cobra.Command, args []string) {
    // In the first for loop, cobra takes i to be the index number of each added
    // "args". arg orders the args in a range, with the first arg first, aso.
    // Spaces delimits args. Add args WITH spaces by closing them in with double
    // quotes.
    // Ex: This will be three args: todo add wash the dishes
    // Ex: This will be one arg: todo add "wash the dishes"
    // Here's how to write:
    // For i and arg, where arg is a range of args:
    // Print i, then ":", then the arg.
    // for i, arg := range args {
    //   fmt.Println(i, ":", arg)
    // }

    // But we rather want to skip the quotes, and can live with adding only one
    // task at a time, so we'll do this:
    task := strings.Join(args, " ")
    fmt.Printf("Added \"%s\" to your task list. \n", task)
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
