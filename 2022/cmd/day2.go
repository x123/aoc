/*
Copyright © 2024 x123
*/
package cmd

import (
	"fmt"
	"strings"
)

var (
	moveMap     = make(map[byte]int)
	strategyMap = make(map[byte]string)
)

func day2_1(content string) {
	moveMap['X'] = 1 // rock
	moveMap['Y'] = 2 // paper
	moveMap['Z'] = 3 // scissors

	fmt.Printf("day:%d part:%d\n", day, part)
	fmt.Printf("content:\n%s\n", content)

	var totalPoints int

	for _, line := range strings.Split(content, "\n") {
		if line != "" {
			moves := strings.Split(line, " ")
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

	moveMap['A'] = 1 // rock
	moveMap['B'] = 2 // paper
	moveMap['C'] = 3 // scissors
	moveMap['X'] = 1 // rock
	moveMap['Y'] = 2 // paper
	moveMap['Z'] = 3 // scissors

	strategyMap['X'] = "lose"
	strategyMap['Y'] = "draw"
	strategyMap['Z'] = "win"

	var totalPoints int
	for _, line := range strings.Split(content, "\n") {
		if line != "" {
			moves := strings.Split(line, " ")
			opponentMove, strategy := moves[0][0], moves[1][0] // offset of 23
			fmt.Printf("opponentMove:%c(%d),strategy:%c(%s)\n", opponentMove, int(opponentMove), strategy, strategyMap[strategy])
			if strategyMap[strategy] == "draw" {
				points := moveMap[opponentMove] + 3
				totalPoints += points
				fmt.Printf("opponentMove:%c(%d),strategy:%c(%s) points:%d\n", opponentMove, int(opponentMove), strategy, strategyMap[strategy], points)
			} else if strategyMap[strategy] == "lose" {
				var myMove byte
				if opponentMove == 'A' {
					myMove = 'C'
				} else if opponentMove == 'B' {
					myMove = 'A'
				} else if opponentMove == 'C' {
					myMove = 'B'
				}
				points := moveMap[myMove]
				totalPoints += points
				fmt.Printf("opponentMove:%c(%d),strategy:%c(%s) points:%d\n", opponentMove, int(opponentMove), strategy, strategyMap[strategy], points)
			} else if strategyMap[strategy] == "win" {
				var myMove byte
				if opponentMove == 'A' {
					myMove = 'B'
				} else if opponentMove == 'B' {
					myMove = 'C'
				} else if opponentMove == 'C' {
					myMove = 'A'
				}
				points := moveMap[myMove] + 6
				totalPoints += points
				fmt.Printf("opponentMove:%c(%d),strategy:%c(%s) points:%d\n", opponentMove, int(opponentMove), strategy, strategyMap[strategy], points)
			}
		}
	}
	fmt.Printf("totalPoints:%d\n", totalPoints)
}
