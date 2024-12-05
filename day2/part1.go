package day2

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com.br/vitorestevam/adventofcode2024/commons"
)

func parseInput(file *os.File) [][]int {
	reports := [][]int{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		report := strings.Split(line, " ")

		fn := func(a string) int {
			v, _ := strconv.Atoi(a)
			return v
		}

		parsedReport := commons.Map(report, fn)

		reports = append(reports, parsedReport)
	}

	return reports
}

func Part1(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	reports := parseInput(file)
	totalSafeReports := 0

	for _, report := range reports {
		oldSign := false
		safe := true
		for i, level := range report {
			if i+1 == len(report) {
				break
			}

			nextLevel := report[i+1]
			diff := level - nextLevel
			abs := math.Abs(float64(diff))
			if abs < 1 || abs > 3 { //rule 2 -> Any two adjacent levels differ by at least one and at most three.
				safe = false
				break
			}

			sign := diff > 0

			if i == 0 {
				oldSign = sign
				continue
			}

			if sign != oldSign { //rule 1 -> The levels are either all increasing or all decreasing.
				safe = false
				break
			}
		}

		if safe {
			totalSafeReports += 1
		}
	}

	return totalSafeReports
}
