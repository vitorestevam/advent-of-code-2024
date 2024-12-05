package day2

import (
	"fmt"
	"math"
	"os"
	"slices"
)

func checkReport(report []int) (safe bool, unsafeIndex int) {
	lenght := len(report)
	safe = true

	fmt.Println("Analizing report", report)
	oldIsIncreasing := false
	for i := 0; i < lenght-1; i += 1 {
		level := report[i]
		nextLevel := report[i+1]

		fmt.Println(i, "checking levels", level, nextLevel)

		diff := level - nextLevel
		abs := math.Abs(float64(diff))
		isIncreasing := nextLevel > level

		if i == 0 {
			oldIsIncreasing = isIncreasing
		}

		if isIncreasing != oldIsIncreasing { //rule 1 -> The levels are either all increasing or all decreasing
			fmt.Println("unsafe on rule 1")
			return false, i
		}

		if abs < 1 || abs > 3 { //rule 2 -> Any two adjacent levels differ by at least one and at most three.
			fmt.Println("unsafe on rule 2")
			return false, i
		}
	}
	return true, 0
}

func Part2(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	reports := parseInput(file)
	totalSafeReports := 0

	for _, report := range reports {
		safe, index := checkReport(report)
		if safe {
			totalSafeReports += 1
			continue
		}

		TrimmedReport := slices.Clone(report)
		TrimmedReport = slices.Delete(TrimmedReport, index, index+1)
		safe, _ = checkReport(TrimmedReport)
		if safe {
			totalSafeReports += 1
			continue
		}

		TrimmedReport = slices.Clone(report)
		TrimmedReport = slices.Delete(TrimmedReport, index+1, index+2)
		safe, _ = checkReport(TrimmedReport)
		if safe {
			totalSafeReports += 1
			continue
		}

		if index > 0 {
			TrimmedReport = slices.Clone(report)
			TrimmedReport = slices.Delete(TrimmedReport, index-1, index)
			safe, _ = checkReport(TrimmedReport)
			if safe {
				totalSafeReports += 1
				continue
			}
		}
	}

	return totalSafeReports
}
