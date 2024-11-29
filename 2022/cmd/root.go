/*
Copyright © 2024 x123
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	day     uint
	part    uint
	preview bool

	rootCmd = &cobra.Command{
		Use:   "2022",
		Short: "aoc 2022",
		Long:  `Advent of Code 2022 in Golang`,
		Run:   runDay,
	}
)

func runDay(cmd *cobra.Command, args []string) {
	thing := [2]uint{day, part}
	switch thing {
	case [2]uint{1, 1}:
		day1_1()
	case [2]uint{1, 2}:
		day1_2()
	default:
		cmd.Help()
	}
}

func loadInput() string {
	var path string
	if preview {
		path = fmt.Sprintf("./%d/preview.txt", day)
	} else {
		path = fmt.Sprintf("./%d/input.txt", day)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func day1_1() {
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

func day1_2() {
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

	// find the sum of the top 3 elves calories
	sort.Ints(calSlice)
	length := len(calSlice)
	top3 := calSlice[length-3:]
	var top3Cals int
	for i, cals := range top3 {
		fmt.Printf("i:%d,cals:%d\n", i, cals)
		top3Cals += cals
	}
	fmt.Printf("top3Cals:%d\n", top3Cals)
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
