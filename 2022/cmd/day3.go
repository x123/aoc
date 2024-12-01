/*
Copyright © 2024 x123
*/
package cmd

import (
	"fmt"
	"slices"
	"strings"
)

func calcPriority(x byte) int {
	y := int(x)
	var priority int
	if y >= 97 && y <= 122 { // 'a' -> 'z"
		priority = y - 96
	} else if y >= 65 && y <= 90 { // 'A' -> 'Z"
		priority = y - 38
	}
	return priority
}

func hashGeneric[T comparable](a []T, b []T) []T {
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

func day3_1(content string) {
	fmt.Printf("day:%d part:%d\n", day, part)
	var totalPriority int

	for _, line := range strings.Split(content, "\n") {
		if line != "" {
			length := len(line)
			fmt.Printf("length:%d,line:%s\n", length, line)
			leftHalf := []byte(line)[:length/2]
			rightHalf := []byte(line)[length/2:]
			fmt.Printf(" leftHalf:%c\n", leftHalf)
			fmt.Printf("rightHalf:%c\n", rightHalf)
			intersect := hashGeneric(leftHalf, rightHalf)[0]
			fmt.Printf("intersect:%c, priority:%d\n", intersect, calcPriority(intersect))
			totalPriority += calcPriority(intersect)
		}
	}
	fmt.Printf("totalPriority:%d\n", totalPriority)
}

func day3_2(content string) {
	for _, line := range strings.Split(content, "\n") {
		if line != "" {
		}
	}
}
