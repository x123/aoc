/*
Copyright © 2024 x123
*/
package cmd

import (
	"fmt"
	// "log"
	// "slices"
	// "sort"
	// "strconv"
	"strings"
)

var (
	moveMap = make(map[byte]int)
)

func day2_1(content string) {
	moveMap['X'] = 1 // rock
	moveMap['Y'] = 2 // paper
	moveMap['Z'] = 3 // scissors

	fmt.Printf("day:%d part:%d\n", day, part)
	fmt.Printf("content:\n%s\n", content)

	var totalPoints int
	for i, line := range strings.Split(content, "\n") {
		fmt.Printf("i:%d,line:'%s'\n", i, line)
		moves := strings.Split(line, " ")
		fmt.Printf("moves:%s\n", moves)
		if line != "" {
			opponentMove, myMove := moves[0][0], moves[1][0] // offset of 23
			fmt.Printf("opponentMove:%c(%d),myMove:%c(%d)\n", opponentMove, int(opponentMove), myMove, int(myMove))

			if int(myMove)-23 == int(opponentMove) {
				// draw
				points := (3 + moveMap[myMove])
				fmt.Printf("DRAW:opponentMove:%c(%d),myMove:%c(%d),points:%d\n", opponentMove, int(opponentMove), myMove, int(myMove), points)
				totalPoints += points
			} else if (int(myMove)-23-int(opponentMove) == 1) || (int(myMove)-23-int(opponentMove) == -2) {
				// win
				points := (6 + moveMap[myMove])
				fmt.Printf("WIN:opponentMove:%c(%d),myMove:%c(%d),points:%d\n", opponentMove, int(opponentMove), myMove, int(myMove), points)
				totalPoints += points
			} else {
				// loss
				points := moveMap[myMove]
				fmt.Printf("LOSS:opponentMove:%c(%d),myMove:%c(%d),points:%d\n", opponentMove, int(opponentMove), myMove, int(myMove), points)
				totalPoints += points
			}
		}
	}
	fmt.Printf("totalPoints:%d\n", totalPoints)
}

func day2_2(content string) {
	fmt.Printf("day:%d part:%d\n", day, part)
	fmt.Printf("content:%s\n", content)
}
