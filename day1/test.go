package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	t.Run("part 1 - test input", func(t *testing.T) {
		expected := 11
		output := Part1("day1/input_test.txt")
		assert.Equal(t, expected, output)
	})
	t.Run("part 1 - real input", func(t *testing.T) {
		expected := 1319616
		output := Part1("day1/input.txt")
		assert.Equal(t, expected, output)
	})
}

func TestPart2(t *testing.T) {

	t.Run("part2 - test input", func(t *testing.T) {
		expected := 31
		output := Part2("day1/input_test.txt")
		assert.Equal(t, expected, output)
	})
	t.Run("part2 - real input", func(t *testing.T) {
		expected := 27267728
		output := Part2("day1/input.txt")
		assert.Equal(t, expected, output)
	})
}
