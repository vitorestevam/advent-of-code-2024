package commons

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateMatrix45(t *testing.T) {
	t.Run("Rotate once", func(t *testing.T) {
		input := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}

		expected := [][]int{
			{1},
			{4, 2},
			{7, 5, 3},
			{8, 6},
			{9},
		}

		result := RotateMatrix45(input)
		assert.Equal(t, expected, result)
	})

	t.Run("rotate string", func(t *testing.T) {
		input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

		trimmedInput := strings.Split(input, "\n")

		fn := func(a string) []string {
			return strings.Split(a, "")
		}
		parsedInput := Map(trimmedInput, fn)
		result := RotateMatrix45(parsedInput)

		fn2 := func(a []string) string {
			return strings.Join(a, "")
		}

		joinedResult := Map(result, fn2)
		parsedResult := strings.Join(joinedResult, "\n")
		fmt.Println(parsedResult)
	})
}

func TestRotateMatrix90(t *testing.T) {
	t.Run("Rotate once", func(t *testing.T) {
		input := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}

		expected := [][]int{
			{7, 4, 1},
			{8, 5, 2},
			{9, 6, 3},
		}

		result := RotateMatrix90(input)
		assert.Equal(t, expected, result)
	})
}
