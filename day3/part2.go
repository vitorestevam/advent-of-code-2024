package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com.br/vitorestevam/adventofcode2024/commons"
)

func parseInput2(fileContent []byte) []string {
	input := string(fileContent)

	var re = regexp.MustCompile(`(?m)(mul\(\d*,\d*\)|do\(\)|don't\(\))`)
	commands := re.FindAllString(input, -1)

	return commands
}

func Part2(filePath string) int {
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	commands := parseInput2(file)

	fmt.Println(commands)

	results := make([]int, len(commands))

	do := true
	for _, command := range commands {
		if strings.Contains(command, "do()") {
			do = true
			continue
		}

		if strings.Contains(command, "don't()") {
			do = false
			continue
		}

		if !do {
			continue
		}

		size := len(command)
		cleannedMul := command[4 : size-1]

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
