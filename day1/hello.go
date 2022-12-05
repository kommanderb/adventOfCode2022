package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"kommanderb/utils"
)

func main() {
	content, err := os.ReadFile("./input.txt")

	utils.CheckError(err)

	strContent := string(content)

	lines := strings.Split(strContent, "\n")

	fmt.Println("Let's begin counting!")

	topThreeMaxCalories := make([]int, 0)
	caloriesAccumulator := 0
	for index := 0; index < len(lines); index++ {
		line := lines[index]

		if len(line) > 0 {
			caloriesAccumulator = accumulateCalories(line, caloriesAccumulator)
		} else {
			topThreeMaxCalories = computeTopThree(topThreeMaxCalories, caloriesAccumulator)
			caloriesAccumulator = 0
		}
	}

	fmt.Println("Most calories found: ", topThreeMaxCalories)
	fmt.Println("Total", totalTopThree(topThreeMaxCalories))
}

func computeTopThree(topThreeMaxCalories []int, caloriesAccumulator int) []int {
	if len(topThreeMaxCalories) < 3 {
		topThreeMaxCalories = append(topThreeMaxCalories, caloriesAccumulator)
	} else {
		for topThreeMaxCaloriesIndex := 0; topThreeMaxCaloriesIndex < len(topThreeMaxCalories); topThreeMaxCaloriesIndex++ {
			minValue, minValueIndex := getMinOfTopThree(topThreeMaxCalories)
			if caloriesAccumulator > minValue {
				topThreeMaxCalories[minValueIndex] = caloriesAccumulator
				break
			}
		}
	}

	fmt.Println("Acc", caloriesAccumulator, "topThree", topThreeMaxCalories)
	return topThreeMaxCalories
}

func accumulateCalories(line string, caloriesAccumulator int) int {
	calories, err := strconv.Atoi(line)
	utils.CheckError(err)
	caloriesAccumulator = caloriesAccumulator + calories

	return caloriesAccumulator
}

func totalTopThree(topThreeMaxCalories []int) int {
	total := 0

	for index := 0; index < len(topThreeMaxCalories); index++ {
		total += topThreeMaxCalories[index]
	}

	return total
}

func getMinOfTopThree(topThreeMaxCalories []int) (int, int) {

	minValue := -1
	minValueIndex := -1

	for index := 0; index < len(topThreeMaxCalories); index++ {
		if minValue < 0 || topThreeMaxCalories[index] < minValue {
			minValue = topThreeMaxCalories[index]
			minValueIndex = index
		}
	}

	return minValue, minValueIndex
}
