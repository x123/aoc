/*
Copyright © 2024 x123
*/
package cmd

import (
	"fmt"
	// "log"
	// "slices"
	// "strconv"
	"strings"
)

func day6LoadMatrix(content string) [][]rune {
	// load matrix
	length := len(strings.Split(content, "\n"))
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
	for _, row := range matrix {
		fmt.Printf("%c\n", row)
	}
	return matrix
}

type position struct {
	y int
	x int
}

func newPerson(y int, x int) *position {
	p := position{y: y, x: x}
	return &p
}

func day6DetermineDirection(playerChar rune) position {
	vector := position{y: 0, x: 0}
	switch playerChar {
	case '^':
		vector = position{y: -1, x: 0}
	case '>':
		vector = position{y: 0, x: 1}
	case 'v':
		vector = position{y: 1, x: 0}
	case '<':
		vector = position{y: 0, x: -1}
	}
	return vector
}

func day6DetermineCharacter(direction position) rune {
	switch direction {
	case position{y: -1, x: 0}:
		return '^'
	case position{y: 0, x: 1}:
		return '>'
	case position{y: 1, x: 0}:
		return 'v'
	case position{y: 0, x: -1}:
		return '<'
	default:
		return 'E' // should never happen
	}
}

func day6Turn(direction position) position {
	switch direction {
	case position{y: -1, x: 0}:
		return position{y: 0, x: 1}
	case position{y: 0, x: 1}:
		return position{y: 1, x: 0}
	case position{y: 1, x: 0}:
		return position{y: 0, x: -1}
	case position{y: 0, x: -1}:
		return position{y: -1, x: 0}
	default:
		return direction
	}
}

func day6FindPlayerPos(matrix [][]rune) position {
	var playerPos position
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '^' || matrix[i][j] == '>' || matrix[i][j] == 'v' || matrix[i][j] == '<' {
				playerPos = position{y: i, x: j}
			}
		}
	}
	return playerPos
}

func day6GetRuneCount(matrix [][]rune, find rune) int {
	var totalCount int
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == find {
				totalCount++
			}
		}
	}
	return totalCount
}

func day6Move(matrix [][]rune) bool {
	var playerPos position
	var moved bool = false

	for i := range matrix {
		if moved {
			break
		}
		for j := range matrix[i] {
			if moved {
				break
			}
			if matrix[i][j] == '^' || matrix[i][j] == '>' || matrix[i][j] == 'v' || matrix[i][j] == '<' {
				playerChar := matrix[i][j]
				playerPos = position{y: i, x: j}
				playerDirection := day6DetermineDirection(matrix[i][j])
				if i+playerDirection.y > len(matrix)-1 || i+playerDirection.y < 0 {
					moved = true
					matrix[i][j] = 'X'
					break
				}
				if j+playerDirection.x > len(matrix[0])-1 || j+playerDirection.x < 0 {
					moved = true
					matrix[i][j] = 'X'
					break
				}
				switch matrix[i+playerDirection.y][j+playerDirection.x] {
				case '.':
					matrix[i+playerDirection.y][j+playerDirection.x] = playerChar
					matrix[i][j] = 'X'
					moved = true
				case 'X':
					matrix[i+playerDirection.y][j+playerDirection.x] = playerChar
					matrix[i][j] = 'X'
					moved = true
				case '#':
					fmt.Printf("oldDirection:%v\n", playerDirection)
					playerDirection = day6Turn(playerDirection)
					fmt.Printf("newDirection:%v\n", playerDirection)
					matrix[i][j] = day6DetermineCharacter(playerDirection)
				}
			}
		}
	}
	fmt.Printf("playerPos.y:%d,x:%d\n", playerPos.y, playerPos.x)
	return moved
}

func day6_1(content string) {
	fmt.Printf("day:%d,part:%d\n", day, part)
	matrix := day6LoadMatrix(content)
	var totalMoves int

	// move until the guard moves away
	for true {
		playerFound := day6FindPlayerPos(matrix)
		if playerFound.x == 0 && playerFound.y == 0 {
			break
		}
		for i := range matrix {
			fmt.Printf("matrix:%c\n", matrix[i])
		}
		moved := day6Move(matrix)
		if moved {
			totalMoves++
		}
	}
	for i := range matrix {
		fmt.Printf("matrix:%c\n", matrix[i])
	}
	fmt.Printf("totalMoves:%d\n", totalMoves)
	distinctVisitedPositions := day6GetRuneCount(matrix, 'X')
	fmt.Printf("distinctVisitedPositions:%d\n", distinctVisitedPositions)
}

func day6_2(content string) {
	fmt.Printf("day:%d,part:%d\n", day, part)
	fmt.Printf("content:%s\n", content)
}
