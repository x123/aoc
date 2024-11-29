/*
Copyright © 2024 x123
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	day uint

	rootCmd = &cobra.Command{
		Use:   "2022",
		Short: "aoc 2022",
		Long:  `Advent of Code 2022 in Golang`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: runDay,
	}
)

func runDay(cmd *cobra.Command, args []string) {
	switch day {
	case 1:
		day1()
	default:
		cmd.Help()
	}
}

func day1() {
	fmt.Printf("day:%d\n", 1)
}

// rootCmd represents the base command when called without any subcommands

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.2022.yaml)")
	rootCmd.PersistentFlags().UintVarP(&day, "day", "d", 0, "day of challenge to run")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
