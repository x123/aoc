/*
Copyright © 2024 x123
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	day     uint
	part    uint
	preview bool

	rootCmd = &cobra.Command{
		Use:   "2024",
		Short: "aoc 2024 ",
		Long:  `Advent of Code 2024 in Golang`,
		Run:   runDay,
	}
)

func runDay(cmd *cobra.Command, args []string) {
	thing := [2]uint{day, part}
	switch thing {
	case [2]uint{1, 1}:
		day1_1(loadInput())
	case [2]uint{1, 2}:
		day1_2(loadInput())
	case [2]uint{2, 1}:
		day2_1(loadInput())
	case [2]uint{2, 2}:
		day2_2(loadInput())
	case [2]uint{3, 1}:
		day3_1(loadInput())
	case [2]uint{3, 2}:
		day3_2(loadInput())
	case [2]uint{4, 1}:
		day4_1(loadInput())
	case [2]uint{4, 2}:
		day4_2(loadInput())
	default:
		cmd.Help()
	}
}

func loadInput() string {
	var path string
	if preview {
		path = fmt.Sprintf("./inputs/%d/preview.txt", day)
	} else {
		path = fmt.Sprintf("./inputs/%d/input.txt", day)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().UintVarP(&day, "day", "d", 0, "day of challenge to run")
	rootCmd.PersistentFlags().UintVarP(&part, "part", "p", 1, "part of challenge to run (1 or 2)")
	rootCmd.PersistentFlags().BoolVar(&preview, "preview", false, "use preview input")
}
