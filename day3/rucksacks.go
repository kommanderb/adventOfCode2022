package main

import (
	"fmt"
	"kommanderb/utils"
	"strings"
)

const PriorityMap = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func splitRucksackCompartments(rucksackContent string) ([]string, int) {
	rucksackCompartmentSize := len(rucksackContent) / 2

	return []string{
		rucksackContent[0:rucksackCompartmentSize],
		rucksackContent[rucksackCompartmentSize:],
	}, rucksackCompartmentSize
}

func getCommonItemType(rucksackCompartments []string, rucksackCompartmentSize int) string {
	var commonItemType = ""
	for index := 0; index < rucksackCompartmentSize; index++ {
		if strings.Contains(rucksackCompartments[1], rucksackCompartments[0][index:index+1]) {
			commonItemType = rucksackCompartments[0][index : index+1]
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
	rucksackCompartments, rucksackCompartmentSize := splitRucksackCompartments(rucksack)

	commonItemType := getCommonItemType(rucksackCompartments, rucksackCompartmentSize)

	return getItemTypePriority(commonItemType)
}

func main() {
	inputFileLines := utils.IngestInputFile("./input.txt")

	var prioritiesAccumulator = 0
	for lineIndex := 0; lineIndex < len(inputFileLines); lineIndex++ {
		if len(inputFileLines[lineIndex]) == 0 {
			continue
		}

		prioritiesAccumulator = prioritiesAccumulator + getRucksackPriorityValue(lineIndex, inputFileLines[lineIndex])
	}

	fmt.Println("Sum of priorities:", prioritiesAccumulator)
}

// Rucksack => 2 compartments
