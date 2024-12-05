package day1

import (
	"bufio"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com.br/vitorestevam/adventofcode2024/commons"
)

func parseInput(file *os.File) ([]int, []int) {
	a, b := []int{}, []int{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		splittedLine := strings.Split(line, "   ")

		{
			v, _ := strconv.Atoi(splittedLine[0])
			a = append(a, v)
		}
		{
			v, _ := strconv.Atoi(splittedLine[1])
			b = append(b, v)
		}
	}

	file.Close()
	return a, b
}

func Part1(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	a, b := parseInput(file)

	slices.Sort(a)
	slices.Sort(b)

	diffs := []int{}

	for i := range a {
		n1 := a[i]
		n2 := b[i]

		diff := n1 - n2
		diff = int(math.Abs(float64(diff)))

		diffs = append(diffs, diff)
	}

	fn := func(a, b int) int {
		return a + b
	}

	sum := commons.Reduce(diffs, fn)

	return sum
}
