package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
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

func main() {
	file, err := os.Open("./input_small.txt")
	if err != nil {
		panic(err)
	}

	a, b := parseInput(file)

	slices.Sort(a)
	slices.Sort(b)

	fmt.Println(a, b)

	diff := []int{}
}
