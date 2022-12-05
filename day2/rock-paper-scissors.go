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
		panic(fmt.Errorf("undefined standard shape! %s", nonStandard))
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

func getShapeToLoseAgainstHandShape(standardHandShape string) string {
	var loosingHandShape = ""

	if standardHandShape == "R" {
		loosingHandShape = "S"
	} else if standardHandShape == "P" {
		loosingHandShape = "R"
	} else if standardHandShape == "S" {
		loosingHandShape = "P"
	}

	return loosingHandShape
}

func getShapeToWinAgainstHandShape(standardHandShape string) string {
	var winningHandShape = ""

	if standardHandShape == "R" {
		winningHandShape = "P"
	} else if standardHandShape == "P" {
		winningHandShape = "S"
	} else if standardHandShape == "S" {
		winningHandShape = "R"
	}

	return winningHandShape
}

func isDraw(standardHandShape string, standardOpponentHandShape string) bool {
	return (standardHandShape == standardOpponentHandShape)
}

func isStandardShape(handShape string) bool {
	return (handShape == "R" || handShape == "P" || handShape == "S")
}

func getScore(opponentHandShape string, handShape string) int {

	if !isStandardShape(opponentHandShape) {
		fmt.Println("opponent shape is not standard")
		opponentHandShape = standardizeHandShape(opponentHandShape)
	}

	if !isStandardShape(handShape) {
		fmt.Println("hand shape is not standard")
		handShape = standardizeHandShape(handShape)
	}

	// fmt.Println(opponentHandShape, "against", handShape)

	var roundScore = 0

	if isDraw(opponentHandShape, handShape) {
		roundScore = 3
	} else if opponentHandShape == getShapeToLoseAgainstHandShape(handShape) {
		roundScore = 6
	}

	return roundScore + getScoreForShape(handShape)
}

func getOutcome(opponentHandShape string, outCome string) int {
	if !isStandardShape(opponentHandShape) {
		opponentHandShape = standardizeHandShape(opponentHandShape)
	}

	var handShape = ""

	if outCome == "X" {
		// lose against handShape
		handShape = getShapeToLoseAgainstHandShape(opponentHandShape)
	} else if outCome == "Y" {
		// draw against handShape
		handShape = opponentHandShape
	} else if outCome == "Z" {
		// win against handShape
		handShape = getShapeToWinAgainstHandShape(opponentHandShape)
	}

	if len(handShape) == 0 {
		panic(fmt.Errorf("unhandled outcome! %s", outCome))
	}

	return getScore(opponentHandShape, handShape)
}

func main() {
	inputLines := utils.IngestInputFile("./input.txt")

	var scoreAccumulator = 0
	for linesIndex := 0; linesIndex < len(inputLines); linesIndex++ {
		if len(inputLines[linesIndex]) == 0 {
			continue
		}

		parts := strings.Split(inputLines[linesIndex], " ")

		// roundScore := getScore(parts[0], parts[1])
		roundScore := getOutcome(parts[0], parts[1])
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
