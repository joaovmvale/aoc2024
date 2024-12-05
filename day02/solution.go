package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const inputFile = "input.txt"

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

func checkIfReportIsSafe(report []int) (bool, int) {
	if len(report) < 2 {
		return true, -1
	}

	reportIsIncreasing := report[1] > report[0]
	for i := 1; i < len(report); i++ {
		difference := report[i] - report[i-1]
		if difference == 0 || difference > 3 || difference < -3 {
			return false, i
		}
		if (reportIsIncreasing && report[i] < report[i-1]) || (!reportIsIncreasing && report[i] > report[i-1]) {
			return false, i
		}
	}
	return true, -1
}

func calculateSafeReportsWithProblemDampener(reports [][]int) int {
	safeReports := 0
	for _, report := range reports {
		reportIsSafe, badLevelIndex := checkIfReportIsSafe(report)
		if reportIsSafe {
			safeReports++
			continue
		}
		if badLevelIndex == -1 {
			continue
		}
		// Remove the bad level and check if the report is safe
		report = append(report[:badLevelIndex], report[badLevelIndex+1:]...)
		reportIsSafe, _ = checkIfReportIsSafe(report)
		if reportIsSafe {
			safeReports++
		}
	}
	return safeReports
}

func calculateSafeReports(reports [][]int) int {
	safeReports := 0
	for _, report := range reports {
		reportIsSafe, _ := checkIfReportIsSafe(report)
		if reportIsSafe {
			safeReports++
			continue
		}
	}
	return safeReports
}

func main() {
	reports := parseInputAndCreateReports(inputFile)
	safeReports := calculateSafeReports(reports)
	fmt.Printf("The amount of safe reports is: %d\n", safeReports)
	safeReportsWithProblemDampener := calculateSafeReportsWithProblemDampener(reports)
	fmt.Printf("The amount of safe reports with the problem dampener is: %d\n", safeReportsWithProblemDampener)
}
