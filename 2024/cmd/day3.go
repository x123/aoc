/*
Copyright © 2024 x123
*/
package cmd

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func day3_1(content string) {
	fmt.Printf("day:%d,part:%d\n", day, part)
	re := regexp.MustCompile(`mul\((?P<x>[0-9]{1,3}),(?P<y>[0-9]{1,3})\)`)
	matches := re.FindAllStringSubmatch(content, -1)
	var totalResult int
	for _, match := range matches {
		fmt.Printf("match:%v\n", match)
		fmt.Printf("match.x:%v\n", match[re.SubexpIndex("x")])
		fmt.Printf("match.y:%v\n", match[re.SubexpIndex("y")])
		x, err := strconv.Atoi(match[re.SubexpIndex("x")])
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(match[re.SubexpIndex("y")])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("result:%d\n", (x * y))
		totalResult += (x * y)
	}
	fmt.Printf("totalResult:%d\n", totalResult)
}

func day3_2(content string) {
	fmt.Printf("day:%d,part:%d\n", day, part)
	length := len(strings.Split(content, "\n"))
	for _, line := range strings.Split(content, "\n")[:length-1] {
		fmt.Printf("%s\n", line)
	}
}
