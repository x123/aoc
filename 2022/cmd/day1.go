/*
Copyright © 2024 x123
*/
package cmd

import (
	"fmt"
	"log"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func day1_1(content string) {
	fmt.Printf("day:%d\n", 1)
	// content := loadInput()

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

func day1_2(content string) {
	fmt.Printf("day:%d\n", 1)
	// content := loadInput()

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
