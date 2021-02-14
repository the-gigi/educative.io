package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"calc-app/pkg/calc"
)

var subtractCmd = &cobra.Command{
	Use:   "subtract",
	Short: "Substract one integer from another",
	Long:  `Substract one integer a from another integer b; result = a -b`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var a, b int
		var err error
		a, err = strconv.Atoi(args[0])
		if err != nil {
			panic("Arguments to `subtract` must be integers")
		}
		b, err = strconv.Atoi(args[1])
		if err != nil {
			panic("Arguments to `subtract` must be integers")
		}

		result := calc.Subtract(a, b, check)
		fmt.Println(result)
	},
}

func init() {
	subtractCmd.Flags().BoolVar(
		&check,
		"check",
		false,
		"check controls if overflow/underflow check is performed")
	rootCmd.AddCommand(subtractCmd)
}
