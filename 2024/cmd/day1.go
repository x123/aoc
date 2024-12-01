/*
Copyright © 2024 x123
*/
package cmd

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func day1_1(content string) {
	fmt.Printf("day:%d,part:%d\n", day, part)

	var leftList []int
	var rightList []int

	var totalDiff int

	for _, line := range strings.Split(content, "\n") {
		if line != "" {
			lineSlice := strings.Split(line, "   ")
			// left, right := lineSlice[0], lineSlice[1]
			left, err := strconv.Atoi(lineSlice[0])
			if err != nil {
				log.Fatal(err)
			}
			right, err := strconv.Atoi(lineSlice[1])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("left:%d,right:%d\n", left, right)
			leftList = append(leftList, left)
			rightList = append(rightList, right)
		}
	}

	sort.Ints(leftList)
	sort.Ints(rightList)
	fmt.Printf("sorted leftList:%v\n", leftList)
	fmt.Printf("sorted rightList:%v\n", rightList)

	for i := range leftList {
		diff := absDiffInt(leftList[i], rightList[i])
		totalDiff += diff
		fmt.Printf("left:%d,right:%d,diff:%d\n", leftList[i], rightList[i], diff)
	}
	fmt.Printf("totalDiff:%d\n", totalDiff)
}

func day1_2(content string) {
	fmt.Printf("day:%d,part:%d\n", day, part)
	fmt.Printf("content:\n%s\n", content)
}
