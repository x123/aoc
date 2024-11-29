/*
Copyright © 2024 x123
*/
package cmd

import (
	// "bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	day uint

	rootCmd = &cobra.Command{
		Use:   "2022",
		Short: "aoc 2022",
		Long:  `Advent of Code 2022 in Golang`,
		Run:   runDay,
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

func loadInput() string {
	path := fmt.Sprintf("./%d/input.txt", day)
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func day1() {
	fmt.Printf("day:%d\n", 1)
	content := loadInput()

	var calorieMap = make(map[int]int)

	elf := 1
	for i, line := range strings.Split(content, "\n") {
		fmt.Printf("i:%d,line:'%s'\n", i, line)
		if line == "" {
			elf += 1
		} else {
			cals, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			calorieMap[elf] += cals
		}
	}

	for i, cals := range calorieMap {
		fmt.Printf("elf:%d,cals:%d\n", i, cals)
	}

	calSlice := make([]int, 0)
	for _, cals := range calorieMap {
		calSlice = append(calSlice, cals)
	}

	max := slices.Max(calSlice)
	fmt.Printf("max:%d\n", max)
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
