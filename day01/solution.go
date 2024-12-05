package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const inputFile = "input.txt"

func parseInputAndSortLists(inputFile string) ([]int, []int) {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}

	var leftList, rightList []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var left, right int
		fmt.Sscanf(scanner.Text(), "%d %d", &left, &right)
		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	return leftList, rightList
}

func calculateDistance(leftList, rightList []int) int {
	distance := 0
	for i := 0; i < len(leftList); i++ {
		if (leftList[i] - rightList[i]) < 0 {
			distance += (rightList[i] - leftList[i])
			continue
		}
		distance += (leftList[i] - rightList[i])
	}

	return distance
}

func calculateSimilarityScore(leftList, rightList []int) int {
	similarityScore := 0
	for i := 0; i < len(leftList); i++ {
		tempScore := 0
		repetitions := 0
		for j := 0; j < len(rightList); j++ {
			if leftList[i] == rightList[j] {
				repetitions++
			}
		}
		tempScore = leftList[i] * repetitions
		similarityScore += tempScore
	}

	return similarityScore
}

func main() {
	leftList, rightList := parseInputAndSortLists(inputFile)
	distance := calculateDistance(leftList, rightList)

	fmt.Printf("The distance is: %d\n", distance)

	similarityScore := calculateSimilarityScore(leftList, rightList)
	fmt.Printf("The similarity score is: %d\n", similarityScore)
}
