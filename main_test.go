package adventofcode2024

import (
	"testing"

	day1 "github.com.br/vitorestevam/adventofcode2024/day1"
	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	t.Run("part 1 - test input", func(t *testing.T) {
		expected := 11
		output := day1.Part1("day1/input_test.txt")
		assert.Equal(t, expected, output)
	})
	t.Run("part 1 - real input", func(t *testing.T) {
		expected := 1319616
		output := day1.Part1("day1/input.txt")
		assert.Equal(t, expected, output)
	})
	t.Run("part2 - test input", func(t *testing.T) {
		expected := 31
		output := day1.Part2("day1/input_test.txt")
		assert.Equal(t, expected, output)
	})
	t.Run("part2 - real input", func(t *testing.T) {
		expected := 27267728
		output := day1.Part2("day1/input.txt")
		assert.Equal(t, expected, output)
	})
}
