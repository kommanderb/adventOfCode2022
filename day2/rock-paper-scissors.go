package main

import (
	"fmt"
	"kommanderb/utils"
	"strings"
)

func standardizeHandShape(nonStandard string) string {
	var standardShape = ""

	if nonStandard == "A" || nonStandard == "X" {
		standardShape = "R"
	} else if nonStandard == "B" || nonStandard == "Y" {
		standardShape = "P"
	} else if nonStandard == "C" || nonStandard == "Z" {
		standardShape = "S"
	}

	if len(standardShape) == 0 {
		panic("Undefined standard shape!")
	}

	return standardShape
}

func getScoreForShape(handShape string) int {
	if handShape == "R" {
		return 1
	} else if handShape == "P" {
		return 2
	} else if handShape == "S" {
		return 3
	}

	fmt.Println("Unhandled hand shape: ", handShape, "Are you trying to play with the pit? Cheater!")

	return 0
}

func getShapeToWinWithHandShape(handShape string) string {
	var loosingHandShape = ""

	if handShape == "R" {
		loosingHandShape = "S"
	} else if handShape == "P" {
		loosingHandShape = "R"
	} else if handShape == "S" {
		loosingHandShape = "P"
	}

	return loosingHandShape
}

func isDraw(standardHandShape string, standardOpponentHandShape string) bool {
	return (standardHandShape == standardOpponentHandShape)
}

func isStandardShape(handShape string) bool {
	return (handShape == "R" || handShape == "P" || handShape == "S")
}

func getScore(handShape string, opponentHandShape string) int {

	if !isStandardShape(handShape) {
		handShape = standardizeHandShape(handShape)
	}

	if !isStandardShape(opponentHandShape) {
		opponentHandShape = standardizeHandShape(opponentHandShape)
	}

	var roundScore = 0

	if isDraw(handShape, opponentHandShape) {
		roundScore = 3
	} else if opponentHandShape == getShapeToWinWithHandShape(handShape) {
		roundScore = 6
	}

	return roundScore + getScoreForShape(handShape)
}

func main() {
	inputLines := utils.IngestInputFile("./test.txt")

	var scoreAccumulator = 0
	for linesIndex := 0; linesIndex < len(inputLines); linesIndex++ {
		if len(inputLines[linesIndex]) == 0 {
			continue
		}

		parts := strings.Split(inputLines[linesIndex], " ")

		roundScore := getScore(parts[1], parts[0])
		fmt.Println("Round #", linesIndex+1, "score:", roundScore)

		scoreAccumulator = scoreAccumulator + roundScore
	}

	fmt.Println("Score:", scoreAccumulator)
}

// A -> Rock -> 1
// B -> Paper -> 2
// C -> Scissors -> 3

// X -> Rock
// Y -> Paper
// Z -> Scissors

// Lose -> 0
// Draw -> 3
// Win -> 6
