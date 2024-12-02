/*
Copyright © 2024 x123
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func day2SplitLine(line string) []int {
	tmpStrings := strings.Split(line, " ")
	// fmt.Printf("%s\n", tmpStrings)
	var result []int
	for _, v := range tmpStrings {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, i)
	}
	return result
}

func day2IsSafe(ints []int) bool {
	safe := true
	asc := false
	desc := false

	lengthInts := len(ints)
	for i := range ints {
		if i == lengthInts-1 {
			return safe
		}
		diff := ints[i+1] - ints[i]
		absDiff := day2AbsDiffInt(ints[i+1], ints[i])

		// set first desc / asc
		if !desc && !asc {
			if diff < 0 {
				desc = true
			}
			if diff > 0 {
				asc = true
			}
		}

		// if we change direction, unsafe
		if diff < 0 && asc {
			safe = false
			return safe
		}
		if diff > 0 && desc {
			safe = false
			return safe
		}

		// if the absolute difference is too steep, unsafe
		if absDiff < 1 || absDiff > 3 {
			safe = false
			return safe
		}
	}
	return safe
}

func day2RemoveElement(slice []int, s int) []int {
	var newSlice []int
	for i := range slice {
		if i != s {
			newSlice = append(newSlice, slice[i])
		}
	}
	return newSlice
}

func day2IsSafePart2(ints []int) bool {
	fmt.Printf("ints:%d\n", ints)
	if day2IsSafe(ints) {
		return true
	} else {
		// remove one element from each position and see if any resuts are safe
		for i := range ints {
			newInts := day2RemoveElement(ints, i)
			fmt.Printf("newInts:%d\n", newInts)
			if day2IsSafe(newInts) {
				return true
			}
		}
	}
	return false
}

func day2AbsDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func day2_1(content string) {
	fmt.Printf("day:%d,part:%d\n", day, part)
	// fmt.Printf("content:\n%s\n", content)
	length := len(strings.Split(content, "\n"))
	var totalSafe int
	for _, line := range strings.Split(content, "\n")[:length-1] {
		ints := day2SplitLine(line)
		fmt.Printf("%d safe:%t\n", ints, day2IsSafe(ints))
		if day2IsSafe(ints) {
			totalSafe += 1
		}
	}
	fmt.Printf("totalSafe:%d\n", totalSafe)
}

func day2_2(content string) {
	fmt.Printf("day:%d,part:%d\n", day, part)
	// fmt.Printf("content:\n%s\n", content)
	length := len(strings.Split(content, "\n"))
	var totalSafe int
	for _, line := range strings.Split(content, "\n")[:length-1] {
		ints := day2SplitLine(line)
		fmt.Printf("%d safe:%t\n", ints, day2IsSafePart2(ints))
		if day2IsSafePart2(ints) {
			totalSafe += 1
		}
	}
	fmt.Printf("totalSafe:%d\n", totalSafe)
}
