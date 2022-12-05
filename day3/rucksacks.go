package main

import (
	"fmt"
	"kommanderb/utils"
	"strings"
)

const PriorityMap = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func splitRucksackCompartments(rucksackContent string) []string {
	rucksackCompartmentSize := len(rucksackContent) / 2

	return []string{
		rucksackContent[0:rucksackCompartmentSize],
		rucksackContent[rucksackCompartmentSize:],
	}
}

func getCommonItemType(rucksackCompartments []string) string {
	var commonItemType = ""
	for index := 0; index < len(rucksackCompartments[0]); index++ {
		search := rucksackCompartments[0][index : index+1]

		foundCount := 0
		for rucksackIndex := 0; rucksackIndex < len(rucksackCompartments); rucksackIndex++ {
			if strings.Contains(rucksackCompartments[rucksackIndex], search) {
				foundCount = foundCount + 1
			} else {
				break
			}
		}

		if foundCount == len(rucksackCompartments) {
			commonItemType = search
		}
	}

	if len(commonItemType) > 0 {
		fmt.Println("Found common item type: ", commonItemType)
	} else {
		panic(fmt.Errorf("common item type not found"))
	}

	return commonItemType
}

func getItemTypePriority(itemType string) int {
	return strings.Index(PriorityMap, itemType) + 1
}

func getRucksackPriorityValue(rucksackNumber int, rucksack string) int {
	rucksackCompartments := splitRucksackCompartments(rucksack)

	commonItemType := getCommonItemType(rucksackCompartments)

	return getItemTypePriority(commonItemType)
}

func orderRucksacks(rucksacks []string) []string {
	permutated := false
	for {
		permutated = false
		for index := 0; index < len(rucksacks); index++ {
			if index+1 < len(rucksacks) {
				if len(rucksacks[index]) > len(rucksacks[index+1]) {
					tempRucksack := rucksacks[index+1]
					rucksacks[index+1] = rucksacks[index]
					rucksacks[index] = tempRucksack
					permutated = true
				}
			} else {
				if len(rucksacks[index]) < len(rucksacks[index-1]) {
					tempRucksack := rucksacks[index-1]
					rucksacks[index-1] = rucksacks[index]
					rucksacks[index] = tempRucksack
					permutated = true
				}
			}
		}

		if !permutated {
			break
		}
	}

	return rucksacks
}

func getRucksacksGroupPriorityValue(rucksacks []string) int {

	orderedRucksacks := orderRucksacks(rucksacks)

	fmt.Println("ordered?", orderedRucksacks)

	commonItemType := getCommonItemType(orderedRucksacks)

	return getItemTypePriority(commonItemType)
}

func main() {
	inputFileLines := utils.IngestInputFile("./input.txt")

	var prioritiesAccumulator = 0
	var rucksacksGroup []string
	for lineIndex := 0; lineIndex < len(inputFileLines); lineIndex++ {
		if len(inputFileLines[lineIndex]) == 0 {
			continue
		}

		rucksacksGroup = append(rucksacksGroup, inputFileLines[lineIndex])

		if len(rucksacksGroup) == 3 {
			prioritiesAccumulator = prioritiesAccumulator + getRucksacksGroupPriorityValue(rucksacksGroup)

			rucksacksGroup = []string{}
		}

	}

	fmt.Println("Sum of priorities:", prioritiesAccumulator)
}
