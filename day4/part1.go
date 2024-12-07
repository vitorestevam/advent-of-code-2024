package day4

import (
	"fmt"
	"os"
	"strings"

	"github.com.br/vitorestevam/adventofcode2024/commons"
)

func inputToMatrix(input string) [][]string {
	splittedInput := strings.Split(input, "\r\n")
	fn := func(a string) []string {
		return strings.Split(a, "")
	}
	parsedInput := commons.Map(splittedInput, fn)

	return parsedInput
}

func matrixToInput(matrix [][]string) string {
	fn := func(a []string) string {
		return strings.Join(a, "")
	}

	joinedMatrix := commons.Map(matrix, fn)
	parsedMatrix := strings.Join(joinedMatrix, "\r\n")
	return parsedMatrix
}

func countXmas(text string) (amount int) {
	// fmt.Println("Counting XMAS on\n", text)

	count := strings.Count(text, "XMAS")
	count += strings.Count(text, "SAMX")
	fmt.Println(count, "Found")
	return count
}

func Part1(filePath string) int {
	byteInput, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	//default 0 degree
	input := string(byteInput)
	count := countXmas(input)

	//rotate 45
	{
		matrixInput := inputToMatrix(input)
		matrixInput = commons.RotateMatrix45(matrixInput)
		input := matrixToInput(matrixInput)
		count += countXmas(input)
	}

	//rotate 90
	{
		matrixInput90 := inputToMatrix(input)
		matrixInput90 = commons.RotateMatrix90(matrixInput90)
		input := matrixToInput(matrixInput90)
		count += countXmas(input)

		{ //rotate 90 + 45 = 135
			matrixInput135 := commons.RotateMatrix45(matrixInput90)
			input := matrixToInput(matrixInput135)
			count += countXmas(input)
		}

		// { //rotate 90+90 = 180
		// 	matrixInput180 := commons.RotateMatrix90(matrixInput90)
		// 	input := matrixToInput(matrixInput180)
		// 	count += countXmas(input)
		// }
	}
	return count
}
