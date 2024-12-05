/*
Copyright © 2024 x123
*/
package cmd

import (
	"fmt"
	// "log"
	// "regexp"
	// "strconv"
	"strings"
)

func day4_1(content string) {
	fmt.Printf("day:%d,part:%d\n", day, part)

	var linesParsed int

	// print lines
	length := len(strings.Split(content, "\n"))
	for _, line := range strings.Split(content, "\n")[:length-1] {
		fmt.Printf("%s\n", line)
		linesParsed++
	}

	// load matrix
	xDim := len(strings.Split(content, "\n")[0])
	yDim := length - 1
	matrix := make([][]rune, yDim)
	for i := range matrix {
		row := strings.Split(content, "\n")[:length-1][i]
		rowArray := strings.Split(row, "")
		var runeArray []rune
		for _, char := range rowArray {
			runeChar := rune(char[0])
			runeArray = append(runeArray, runeChar)
		}
		matrix[i] = runeArray
	}

	fmt.Printf("xDim:%d,xYdim:%d\n", xDim, yDim)
	// fmt.Printf("matrix:%c\n", matrix)
	for _, row := range matrix {
		fmt.Printf("%c\n", row)
	}
	// result := day4FindChars(matrix, 'X')
	// for i := range result {
	// 	fmt.Printf("result:%v\n", result[i])
	// }
	day4FindChars(matrix)
}

func day4FindChars(matrix [][]rune) {
	fmt.Printf("len(matrix):%d\n", len(matrix))
	var totalFound int
	found := make([][]bool, len(matrix))
	for i := range found {
		found[i] = make([]bool, len(matrix[0]))
	}
	for y, line := range matrix {
		for x := range line {
			runeChar := matrix[y][x]
			if runeChar == 'X' || runeChar == 'S' {
				// fmt.Printf("runeChar:%c,pos:y%d,x%d\n", runeChar, y, x)
				strLine := string(line[x : x+4])

				// horizontal
				if strLine == "XMAS" || strLine == "SAMX" {
					for i := x; i < x+4; i++ {
						found[y][i] = true
					}
					totalFound++
					fmt.Printf("strLine:%s\n", strLine)
				}

				// vertical
				if y+4 <= len(matrix) {
					var vert []rune = []rune{matrix[y][x]}
					for i := y + 1; i < y+4; i++ {
						vert = append(vert, matrix[i][x])
					}
					strVert := string(vert)
					if strVert == "XMAS" || strVert == "SAMX" {
						for i := y; i < y+4; i++ {
							found[i][x] = true
						}
						fmt.Printf("strVert:%s\n", strVert)
						totalFound++
					}
				}

				// diagonal right
				if y+4 <= len(matrix) {
					if x+4 <= len(matrix[0]) {
						var diagRight []rune = []rune{matrix[y][x]}
						for i := 1; i < 4; i++ {
							diagRight = append(diagRight, matrix[y+i][x+i])
						}
						strDiagRight := string(diagRight)
						if strDiagRight == "XMAS" || strDiagRight == "SAMX" {
							for i := 0; i < 4; i++ {
								found[y+i][x+i] = true
							}
							fmt.Printf("strDiagRight:%s\n", strDiagRight)
							totalFound++
						}

					}
				}
				// diagonal left
				if y+4 <= len(matrix) {
					if x-3 >= 0 {
						var diagLeft []rune = []rune{matrix[y][x]}
						for i := 1; i < 4; i++ {
							diagLeft = append(diagLeft, matrix[y+i][x-i])
						}
						strDiagLeft := string(diagLeft)
						if strDiagLeft == "XMAS" || strDiagLeft == "SAMX" {
							for i := 0; i < 4; i++ {
								found[y+i][x-i] = true
							}
							fmt.Printf("strDiagLeft:%s\n", strDiagLeft)
							totalFound++
						}

					}
				}
			}
		}
	}

	// pretty print matrix
	for y := range len(matrix) {
		for x := range len(matrix[y]) {
			if found[y][x] {
				fmt.Printf("%c", matrix[y][x])
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("totalFound:%d\n", totalFound)
}

func day4_2(content string) {
	fmt.Printf("day:%d,part:%d\n", day, part)
	length := len(strings.Split(content, "\n"))
	var totalSafe int
	for _, line := range strings.Split(content, "\n")[:length-1] {
		ints := day2SplitLine(line)
		fmt.Printf("%d safe:%t\n", ints, day2IsSafe(ints))
		if day2IsSafe(ints) {
			totalSafe += 1
		}
	}
}
