/*
Copyright © 2024 x123
*/
package cmd

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

func day5LoadRulesPages(content string) ([][2]int, [][]int) {
	rules := make([][2]int, 0)
	length := len(strings.Split(content, "\n"))
	pages := make([][]int, 0)
	foundPages := false
	for _, line := range strings.Split(content, "\n")[:length-1] {
		// rules
		if !foundPages && line != "" {
			strArray := strings.Split(line, "|")
			var rulesArray [2]int
			for i := range strArray {
				x, err := strconv.Atoi(strArray[i])
				if err != nil {
					log.Fatal(err)
				}
				rulesArray[i] = x
			}
			rules = append(rules, rulesArray)
		} else if foundPages && line != "" {
			// pages
			strArray := strings.Split(line, ",")
			var pagesArray []int
			for i := range strArray {
				x, err := strconv.Atoi(strArray[i])
				if err != nil {
					log.Fatal(err)
				}
				pagesArray = append(pagesArray, x)
			}
			pages = append(pages, pagesArray)
		}
		if line == "" {
			foundPages = true
		}
	}
	return rules, pages
}

// find the intersect of two arrays
func day5HashGeneric[T comparable](a []T, b []T) []T {
	set := make([]T, 0)
	hash := make(map[T]struct{})

	for _, v := range a {
		hash[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			if !slices.Contains(set, v) {
				set = append(set, v)
			}
		}
	}

	return set
}

// find the position of b in a
func day5FindPos(a []int, b int) int {
	for i := range a {
		if a[i] == b {
			return i
		}
	}
	return -1
}

// find the value at the middle position of an array
func day5FindMiddle(a []int) int {
	if len(a)%2 == 0 {
		fmt.Printf("day5FindMiddle must accept an odd length array\n")
		return -1
	}
	midPos := len(a) / 2
	return a[midPos]
}

func day5_1(content string) {
	fmt.Printf("day:%d,part:%d\n", day, part)
	var totalMiddleSum int

	rules, pages := day5LoadRulesPages(content)
	fmt.Printf("rules:%v\n", rules)
	fmt.Printf("pages:%v\n", pages)

	for _, pageSet := range pages {
		valid := true
		for _, rulePair := range rules {
			left := rulePair[0]
			right := rulePair[1]
			intersect := day5HashGeneric(pageSet, []int{left, right})
			if len(intersect) == 2 {
				leftPos := day5FindPos(pageSet, left)
				rightPos := day5FindPos(pageSet, right)
				if leftPos == -1 || rightPos == -1 {
					log.Fatal("day5FindPos got -1")
				}
				if leftPos > rightPos {
					valid = false
				}
			}
		}
		if valid {
			fmt.Printf("valid pageset:%v\n", pageSet)
			middle := day5FindMiddle(pageSet)
			fmt.Printf("middle:%v\n", middle)
			totalMiddleSum += middle
		}
	}
	fmt.Printf("totalMiddleSum:%d\n", totalMiddleSum)
}

func day5_2(content string) {
	fmt.Printf("day:%d,part:%d\n", day, part)
	var totalMiddleSum int

	rules, pages := day5LoadRulesPages(content)
	fmt.Printf("rules:%v\n", rules)
	fmt.Printf("pages:%v\n", pages)

	for _, pageSet := range pages {
		valid := true
		for _, rulePair := range rules {
			left := rulePair[0]
			right := rulePair[1]
			intersect := day5HashGeneric(pageSet, []int{left, right})
			if len(intersect) == 2 {
				leftPos := day5FindPos(pageSet, left)
				rightPos := day5FindPos(pageSet, right)
				if leftPos == -1 || rightPos == -1 {
					log.Fatal("day5FindPos got -1")
				}
				if leftPos > rightPos {
					valid = false
				}
			}
		}
		if !valid {
			fmt.Printf("invalid pageset:%v\n", pageSet)
			var applicableRules [][2]int
			for _, rulePair := range rules {
				left := rulePair[0]
				right := rulePair[1]
				intersect := day5HashGeneric(pageSet, []int{left, right})
				if len(intersect) == 2 {
					applicableRules = append(applicableRules, [2]int{left, right})
					leftPos := day5FindPos(pageSet, left)
					rightPos := day5FindPos(pageSet, right)
					if leftPos == -1 || rightPos == -1 {
						log.Fatal("day5FindPos got -1")
					}
					if leftPos > rightPos {
						fmt.Printf("pageSet.left:%d,pageSet.right:%d,not respecting rulePair:%d\n", pageSet[leftPos], pageSet[rightPos], rulePair)
						valid = false
					}
				}
			}

			// recursively fix the pageset until it is valid based on the applicable rules
			changed := true
			for changed {
				changed = fixPageSet(pageSet, applicableRules)
			}

			fmt.Printf("fixed pageset:%v\n", pageSet)
			fmt.Printf("applicableRules:%d\n", applicableRules)
			middle := day5FindMiddle(pageSet)
			fmt.Printf("middle:%v\n", middle)
			totalMiddleSum += middle
		}
	}
	fmt.Printf("totalMiddleSum:%d\n", totalMiddleSum)
}

// reorder pageSet on the first match by swapping, designed to be used recursively
func fixPageSet(pageSet []int, applicableRules [][2]int) bool {
	for _, rulePair := range applicableRules {
		left := rulePair[0]
		right := rulePair[1]
		leftPos := day5FindPos(pageSet, left)
		rightPos := day5FindPos(pageSet, right)
		if leftPos > rightPos {
			var tmp = slices.Clone(pageSet)
			pageSet[rightPos] = tmp[leftPos]
			pageSet[leftPos] = tmp[rightPos]
			return true
		}
	}
	return false
}
