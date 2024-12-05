package day1

import (
	"os"
	"slices"

	"github.com.br/vitorestevam/adventofcode2024/commons"
)

func Part2(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	a, b := parseInput(file)

	//optional, but faster
	slices.Sort(a)
	slices.Sort(b)

	simillarityScores := []int{}

	for _, i := range a {
		count := 0

		for _, j := range b {
			if i == j {
				count += 1
				continue
			}
			if i < j {
				break
			}
		}

		simillarity := count * i

		simillarityScores = append(simillarityScores, simillarity)
	}

	fn := func(a, b int) int {
		return a + b
	}

	sum := commons.Reduce(simillarityScores, fn)

	return sum
}
