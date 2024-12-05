package day3

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com.br/vitorestevam/adventofcode2024/commons"
)

func parseInput(fileContent []byte) []string {
	input := string(fileContent)

	var re = regexp.MustCompile(`(?m)(mul\(\d*,\d*\))`)
	muls := re.FindAllString(input, -1)

	return muls
}

func Part1(filePath string) int {
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	muls := parseInput(file)

	results := make([]int, len(muls))

	for _, mul := range muls {
		size := len(mul)
		cleannedMul := mul[4 : size-1]

		numbers := strings.Split(cleannedMul, ",")
		n1, _ := strconv.Atoi(numbers[0])
		n2, _ := strconv.Atoi(numbers[1])
		result := n1 * n2

		results = append(results, result)
	}

	fn := func(a, b int) int {
		return a + b
	}

	sum := commons.Reduce(results, fn)

	return sum
}
