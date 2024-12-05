package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const inputFile = "input.txt"

// Parses the input file and creates a list of reports.
func parseInputAndCreateReports(inputFile string) [][]int {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var reports [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var report []int
		for _, level := range strings.Fields(scanner.Text()) {
			var tempLevel int
			fmt.Sscanf(level, "%d", &tempLevel)
			report = append(report, tempLevel)
		}
		reports = append(reports, report)
	}
	return reports
}

// Checks if a report is safe based on level differences.
func isSafe(report []int) bool {
	if len(report) < 2 {
		return true
	}

	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(report); i++ {
		difference := report[i] - report[i-1]
		if difference <= 0 || difference > 3 {
			isIncreasing = false
		}
		if difference >= 0 || difference < -3 {
			isDecreasing = false
		}
	}

	return isIncreasing || isDecreasing
}

// Calculates the number of safe reports using the Problem Dampener.
func calculateSafeReportsWithProblemDampener(reports [][]int) int {
	safeReports := 0
	for _, report := range reports {
		// Check if the report is already safe.
		if isSafe(report) {
			safeReports++
			continue
		}

		// Try removing each level and recheck safety.
		for i := 0; i < len(report); i++ {
			newReport := append([]int{}, report[:i]...)
			newReport = append(newReport, report[i+1:]...)
			if isSafe(newReport) {
				safeReports++
				break
			}
		}
	}
	return safeReports
}

// Calculates the number of safe reports without the Problem Dampener.
func calculateSafeReports(reports [][]int) int {
	safeReports := 0
	for _, report := range reports {
		if isSafe(report) {
			safeReports++
		}
	}
	return safeReports
}

func main() {
	reports := parseInputAndCreateReports(inputFile)

	// Safe reports without the Problem Dampener.
	safeReports := calculateSafeReports(reports)
	fmt.Printf("The amount of safe reports is: %d\n", safeReports)

	// Safe reports with the Problem Dampener.
	safeReportsWithProblemDampener := calculateSafeReportsWithProblemDampener(reports)
	fmt.Printf("The amount of safe reports with the problem dampener is: %d\n", safeReportsWithProblemDampener)
}
