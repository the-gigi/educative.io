package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"calc-app/pkg/calc"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add two integers",
	Long:  `Add two integers a and b; result = a + b`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var a, b int
		var err error
		a, err = strconv.Atoi(args[0])
		if err != nil {
			panic("Arguments to `add` must be integers")
		}
		b, err = strconv.Atoi(args[1])
		if err != nil {
			panic("Arguments to `add` must be integers")
		}

		result := calc.Add(a, b, check)
		fmt.Println(result)
	},
}

func init() {
	addCmd.Flags().BoolVar(
		&check,
		"check",
		false,
		"check controls if overflow/underflow check is performed")
	rootCmd.AddCommand(addCmd)
}
