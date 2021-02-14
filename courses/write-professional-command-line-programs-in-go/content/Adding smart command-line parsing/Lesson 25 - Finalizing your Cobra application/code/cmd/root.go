/*
Copyright © 2020

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var check bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "calc-app",
	Short: "Calculate arithmetic expressions",
	Long: `Calculate arithmetic expressions.
           
           It can add integers and subtract integers`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
