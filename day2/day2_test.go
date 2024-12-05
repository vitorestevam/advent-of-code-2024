package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	t.Run("part 1 - test input", func(t *testing.T) {
		expected := 11
		output := Part1("input_test.txt")
		assert.Equal(t, expected, output)
	})
	t.Run("part 1 - real input", func(t *testing.T) {
		expected := 1319616
		output := Part1("input.txt")
		assert.Equal(t, expected, output)
	})
}

func TestPart2(t *testing.T) {
	t.Run("Test input", func(t *testing.T) {
		expected := 4
		output := Part2("input_test.txt")
		assert.Equal(t, expected, output)
	})
	t.Run("Real input", func(t *testing.T) {
		expected := 634
		output := Part2("input.txt")
		assert.Equal(t, expected, output)
	})
}
