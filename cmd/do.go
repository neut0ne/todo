package cmd

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete.",
	// To be able to run "do task 1 2 3", we need to convert all task strings into
	// integers. Create a var called ids and make it an integer slice.
	// for all index integers, arg is a range of args where:
	// we do a string conversion of A to i. If it doesn't work, throw an error;
	// if it works, append id to ids.
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		fmt.Println(ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
