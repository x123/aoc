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

func day4_1(content string) {
	fmt.Printf("day:%d part:%d\n", day, part)

	var min int
	var max int
	var totalContains int

	length := len(strings.Split(content, "\n"))
	for _, line := range strings.Split(content, "\n")[:length-1] {
		fmt.Println(line)
		sliced := strings.Split(line, ",")
		leftStart, err := strconv.Atoi(strings.Split(sliced[0], "-")[0])
		if err != nil {
			log.Fatal(err)
		}
		leftEnd, err := strconv.Atoi(strings.Split(sliced[0], "-")[1])
		if err != nil {
			log.Fatal(err)
		}
		rightStart, err := strconv.Atoi(strings.Split(sliced[1], "-")[0])
		if err != nil {
			log.Fatal(err)
		}
		rightEnd, err := strconv.Atoi(strings.Split(sliced[1], "-")[1])
		if err != nil {
			log.Fatal(err)
		}
		if leftStart < min {
			min = leftStart
		}
		if rightStart < min {
			min = rightStart
		}
		if leftEnd > max {
			max = leftEnd
		}
		if rightEnd > max {
			max = rightEnd
		}
		var leftSlice []int
		for i := leftStart; i <= leftEnd; i++ {
			leftSlice = append(leftSlice, i)
		}
		var rightSlice []int
		for i := rightStart; i <= rightEnd; i++ {
			rightSlice = append(rightSlice, i)
		}
		fmt.Printf(" leftSlice:%d\n", leftSlice)
		fmt.Printf("rightSlice:%d\n", rightSlice)
		rightContainsLeft := true
		for i := leftStart; i <= leftEnd; i++ {
			if !slices.Contains(rightSlice, i) {
				rightContainsLeft = false
			}
		}
		leftContainsRight := true
		for i := rightStart; i <= rightEnd; i++ {
			if !slices.Contains(leftSlice, i) {
				leftContainsRight = false
			}
		}
		if leftContainsRight || rightContainsLeft {
			totalContains += 1
			fmt.Printf("leftContainsRight\n")
		}
	}
	fmt.Printf("totalContains:%d\n", totalContains)
}

func day4_2(content string) {
	fmt.Printf("day:%d part:%d\n", day, part)
	var min int
	var max int
	var totalContains int

	length := len(strings.Split(content, "\n"))
	for _, line := range strings.Split(content, "\n")[:length-1] {
		fmt.Println(line)
		sliced := strings.Split(line, ",")
		leftStart, err := strconv.Atoi(strings.Split(sliced[0], "-")[0])
		if err != nil {
			log.Fatal(err)
		}
		leftEnd, err := strconv.Atoi(strings.Split(sliced[0], "-")[1])
		if err != nil {
			log.Fatal(err)
		}
		rightStart, err := strconv.Atoi(strings.Split(sliced[1], "-")[0])
		if err != nil {
			log.Fatal(err)
		}
		rightEnd, err := strconv.Atoi(strings.Split(sliced[1], "-")[1])
		if err != nil {
			log.Fatal(err)
		}
		if leftStart < min {
			min = leftStart
		}
		if rightStart < min {
			min = rightStart
		}
		if leftEnd > max {
			max = leftEnd
		}
		if rightEnd > max {
			max = rightEnd
		}
		var leftSlice []int
		for i := leftStart; i <= leftEnd; i++ {
			leftSlice = append(leftSlice, i)
		}
		var rightSlice []int
		for i := rightStart; i <= rightEnd; i++ {
			rightSlice = append(rightSlice, i)
		}
		fmt.Printf(" leftSlice:%d\n", leftSlice)
		fmt.Printf("rightSlice:%d\n", rightSlice)
		rightContainsLeft := false
		for i := leftStart; i <= leftEnd; i++ {
			if slices.Contains(rightSlice, i) {
				rightContainsLeft = true
			}
		}
		leftContainsRight := false
		for i := rightStart; i <= rightEnd; i++ {
			if slices.Contains(leftSlice, i) {
				leftContainsRight = true
			}
		}
		if leftContainsRight || rightContainsLeft {
			totalContains += 1
			fmt.Printf("leftContainsRight\n")
		}
	}
	fmt.Printf("totalContains:%d\n", totalContains)
}
