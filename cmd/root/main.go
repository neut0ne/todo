package main

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
  Use:   "Todo",
  Short: "Todo is a small and efficient task manager.",
  Long: `A Fast and intuitive basic task mamagement tool built with
                love by neut0ne and friends in Go.
                Complete documentation is available at http://todo.neut0ne.go`,
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
  },
}
